package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

func transpile(code string) (string, []string) {
	client := getClient()
	_ = client

	msg, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: 2048,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(code)),
		},
		System: []anthropic.TextBlockParam{
			{Text: systemTxt},
		},
		Model: "claude-haiku-4-5",
		Tools: getTools(),
	})
	if err != nil {
		log.Fatal(err)
	}

	type NewFile struct {
		name string
		content string
	}

	dependencies := []string{}
	var files []NewFile = []NewFile{}

	for _, block := range msg.Content {
		if toolUse, ok := block.AsAny().(anthropic.ToolUseBlock); ok {
			_ = toolUse
			
			switch toolUse.Name {
			case "install":
				var input struct {
					Name string `json:"name"`
				}

				if err := json.Unmarshal(toolUse.Input, &input); err != nil {
					log.Printf("Error while parsing a tool usage. Continuing")
					continue
				}

				dependencies = append(dependencies, input.Name)

			case "new-file":
				var input struct {
					Name string `json:"name"`
					Content string `json:"content"`
				}

				if err := json.Unmarshal(toolUse.Input, &input); err != nil {
					log.Println("Error while parsing a tool usage")
					continue
				}

				files = append(files, NewFile{
					name: input.Name,
					content: input.Content,
				})
			}
		}
	}

	return msg.Content[0].Text, dependencies
}
