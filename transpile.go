package main

import (
	"context"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

func transpile(code string) string {
	client := getClient()
	_ = client

	msg, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(code)),
		},
		System: []anthropic.TextBlockParam{
			{Text: systemTxt},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	return msg.Content[0].Text
}
