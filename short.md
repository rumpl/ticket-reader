---
author: David Gageot & Djordje Lukić
paging: Slide %d / %d
theme: ./theme.json
---

# Docker Agent, Docker's Agentic Swiss Army Knife.

AI is revolutionizing our development-related jobs.

At Docker, we're playing the role of scouts and completely changing
the way we work.

To help us, we've created `Docker Agent`, an agentic framework for building agents.

Code agents, but not only.

`Open-Source, Model-Agnostic, often without a single line of code.`

---

# Who are we?

**Djordje Lukić** - Principal Engineer at **Docker**

**David Gageot** - Senior Principal Engineer at **Docker**

## AI products for the past year and a half:

+ `Gordon`
+ `MCP Gateway`
+ `MCP Toolkit`
+ `Docker Agent`
+ And more 🐳

---

# AI for the past year and a half

It feels like an eternity!

+ Started with Gordon using `llama3` then `gpt-4o`
+ No `MCP`
+ No `Opus`
+ Early days of `Copilot`
+ Chat bots but no agents
+ Coding with AI was possible but a bumpy road
+ The beginnings of *vibe-coding*

---

# And now?

A brand new world:

+ `Opus` ~~4.6~~`4.7`, `GPT 5.4`
+ `MCP` is already almost out
+ Agent **orchestrators**
+ Codebases `100%` coded by agents

---

# What about you?

+ Who uses AI to code?

---

# What about you?

+ Who uses AI to code?
+ Who doesn't code anymore?

---

# What about you?

+ Who uses AI to code?
+ Who doesn't code anymore?
+ Who builds agents for something other than coding?

---

# Docker Agent

## A true Swiss Army Knife
+ Open-Source
+ Open: `OpenAI`, `Anthropic`, `Gemini`, `Mistral`, `Docker Model Runner`, `Ollama` ...

## Rich
+ `Docker Agent` is mainly coded with `Docker Agent`
+ Lots and lots of features!

## Usable in multiple ways
+ `YAML`
+ Go `SDK`
+ Used by `Docker` in production for `Gordon`

---

# First Demo

`Docker Agent` can be used to define any kind of agent.

**From the most useful to the least useful.**

## A totally useless example:

```bash
$ docker agent run "pirate.yaml"
```

---

# Demo: Pirate

## Just YAML and nothing else

```yaml
agents:
  root:
    model: anthropic/claude-sonnet-4-5
    instruction: Always answer by talking like a pirate.
    welcome_message: |
      Ahoy! I be yer pirate guide, ready to set sail on the seas o' knowledge!
      What be yer quest? 🏴‍☠️
```

---

# Demo - A specialized agent

**Can we write highly specialized agents?**

+ Able to use *cheaper* models
+ And *faster* ones
+ Thus with a *wider choice* of models
+ With the minimum set of tools, to limit the impact of potential errors

---

# Demo - A specialized agent

There are `APIs` all around us.

*Can we write an agent that uses these APIs as tools?*

It's easy to turn an Open API-compatible API into a toolbox.

```yaml
- type: openapi
  url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

**Here it's a Pokemon API, elsewhere it could be a business API.**

---

# Demo - A Pokemon-specialized agent

## The YAML

```yaml
agents:
  root:
    model: anthropic/claude-haiku-4-5
    instruction: |
      You are a Pokémon expert.
      Be brief in your answers.
      Be fun too!
      IMPORTANT: Always use the French names of Pokémon.
    toolsets:
      - type: openapi
        url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

## Running the agent

```bash
$ docker agent run "./pokemon.yaml"
```

---

# Demo - A Pokemon-specialized agent

## The YAML

```yaml
agents:
  root:
    model: anthropic/claude-haiku-4-5
    instruction: |
      You are a Pokémon expert.
      Be brief in your answers.
      Be fun too!
      IMPORTANT: Always use the French names of Pokémon.
    toolsets:
      - type: openapi
        url: "https://raw.githubusercontent.com/PokeAPI/pokeapi/master/openapi.yml"
```

## Running the agent

```bash
$ docker agent run "./pokemon.yaml"
```

## Or a more advanced version:

```bash
$ docker agent run "./pokemon-plus.yaml"
```

---

# Demo - How do we optimize this agent?

We write `evals` and check whether our agent gets a good score.

## Running the evals

```bash
$ docker agent eval "./pokemon-plus.yaml"
```

## And then?

We tweak the `YAML` by hand and loop!

---

# Demo - How do we optimize this agent?

Much better: use another agent to optimize our agent.

## Automatic optimization

```bash
$ docker agent run "./eval-expert.yaml" "Optimize pokemon-plus.yaml"
```

---

# What we didn't show you

+ Code Agents
+ Go SDK
+ Structured Output
+ A2A
+ ACP
+ MCP (local, remote, oauth)
+ Handoff
+ Alloy
+ RAG
+ Code Mode
+ Slash commands
+ Hooks
+ Sandbox
+ Routing Mode
+ More tools (Http, TODO, Memory, Scripts, LSP, Background Agents...)
+ ...

---

# Questions

Thanks everyone!
