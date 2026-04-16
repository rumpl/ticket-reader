package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/docker/docker-agent/pkg/agent"
	"github.com/docker/docker-agent/pkg/chat"
	"github.com/docker/docker-agent/pkg/config/latest"
	"github.com/docker/docker-agent/pkg/environment"
	"github.com/docker/docker-agent/pkg/model/provider/openai"
	"github.com/docker/docker-agent/pkg/model/provider/options"
	"github.com/docker/docker-agent/pkg/runtime"
	"github.com/docker/docker-agent/pkg/session"
	"github.com/docker/docker-agent/pkg/team"
	"github.com/google/jsonschema-go/jsonschema"
)

type Ticket struct {
	Price float64 `json:"price" jsonschema:"The total price"`
	Store string  `json:"store" jsonschema:"The store name"`
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelError)

	if err := run(context.Background()); err != nil {
		log.Println(err)
	}
}

func run(ctx context.Context) error {
	schema, err := jsonschema.For[Ticket](&jsonschema.ForOptions{})
	if err != nil {
		return err
	}

	llm, err := openai.NewClient(
		ctx,
		&latest.ModelConfig{
			Provider: "openai",
			Model:    "gpt-5.4",
		},
		environment.NewDefaultProvider(),
		options.WithStructuredOutput(&latest.StructuredOutput{
			Name:   "ticket",
			Schema: SchemaToMap(schema),
		}),
	)
	if err != nil {
		return err
	}

	human := agent.New(
		"root",
		"Your job is to read a receipt and extract the total price from it. You will be given the receipt as text. You should only return the total price in a structured format.",
		agent.WithModel(llm),
	)

	humanTeam := team.New(team.WithAgents(human))

	rt, err := runtime.New(humanTeam)
	if err != nil {
		return err
	}

	data, err := imageData(os.Args[1])
	if err != nil {
		return err
	}
	multiContent := []chat.MessagePart{
		{
			Type: chat.MessagePartTypeText,
			Text: "Give me the total price",
		},
		{
			Type: chat.MessagePartTypeImageURL,
			ImageURL: &chat.MessageImageURL{
				URL:    data,
				Detail: chat.ImageURLDetailAuto,
			},
		},
	}

	um := session.UserMessage("Give me the total price", multiContent...)
	sess := session.New()
	sess.AddMessage(um)

	messages, err := rt.Run(ctx, sess)
	if err != nil {
		return err
	}

	fmt.Println(messages[len(messages)-1].Message.Content)
	return nil
}
