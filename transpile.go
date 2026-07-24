package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/anthropics/anthropic-sdk-go"
)

type NewFile struct {
	name string
	content string
}

func transpile(code string) ([]NewFile, []string) {
	client := getClient()

	messages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock(code)),
	}

	dependencies := []string{}
	var files []NewFile = []NewFile{}

	for {
		msg, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
			MaxTokens: 4096,
			Messages:  messages,
			System: []anthropic.TextBlockParam{
				{Text: systemTxt},
			},
			Model: "claude-haiku-4-5",
			Tools: getTools(),
		})
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, msg.ToParam())

		var toolResults []anthropic.ContentBlockParamUnion
		sawToolUse := false

		for _, block := range msg.Content {
			if toolUse, ok := block.AsAny().(anthropic.ToolUseBlock); ok {
				sawToolUse = true
				switch toolUse.Name {
				case "install":
					var input struct {
						Name string `json:"name"`
					}
					if err := json.Unmarshal(toolUse.Input, &input); err != nil {
						log.Printf("Error parsing install input, continuing")
					} else {
						dependencies = append(dependencies, input.Name)
					}
					toolResults = append(toolResults, anthropic.NewToolResultBlock(toolUse.ID, "installed", false))

				case "new-file":
					var input struct {
						Name    string `json:"name"`
						Content string `json:"content"`
					}
					if err := json.Unmarshal(toolUse.Input, &input); err != nil {
						log.Println("Error parsing new-file input")
					} else {
						files = append(files, NewFile{name: input.Name, content: input.Content})
					}
					toolResults = append(toolResults, anthropic.NewToolResultBlock(toolUse.ID, "file created", false))
				}
			}
		}

		if !sawToolUse || msg.StopReason != "tool_use" {
			break
		}

		messages = append(messages, anthropic.NewUserMessage(toolResults...))
	}

	return files, dependencies
}