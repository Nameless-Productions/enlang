package main

import (
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func getClient() anthropic.Client {
	return anthropic.NewClient(
		option.WithAPIKey(os.Getenv("CLAUDE")),
	)
}