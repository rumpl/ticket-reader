package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/docker/docker-agent/pkg/chat"
)

func imageData(path string) (string, error) {
	imgData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	mimeType := chat.DetectMimeType(path)

	resized, err := chat.ResizeImage(imgData, mimeType)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("data:%s;base64,%s", resized.MimeType, base64.StdEncoding.EncodeToString(resized.Data)), nil
}
