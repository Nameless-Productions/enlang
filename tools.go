package main

import "github.com/anthropics/anthropic-sdk-go"

func getTools() []anthropic.ToolParam {
	return []anthropic.ToolParam{
		{
			Name: "install",
			Description: anthropic.String("Install a go dependencie (a go get command)"),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"name": map[string]any{
						"type": "string",
						"description": "the name of the dependencie you want to get",
					},
				},
			},
		},
	}
}