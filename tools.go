package main

import "github.com/anthropics/anthropic-sdk-go"

func getTools() []anthropic.ToolUnionParam {
	toolParams := []anthropic.ToolParam{
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
		{
			Name: "new-file",
			Description: anthropic.String("Create a new file"),
			InputSchema: anthropic.ToolInputSchemaParam{
				Properties: map[string]any{
					"name": map[string]any{
						"type": "string",
						"description": "The name of the file (ex.: main.go)",
					},
					"content": map[string]any{
						"type": "string",
						"description": "The Go code of the file",
					},
				},
			},
		},
	}

	tools := make([]anthropic.ToolUnionParam, len(toolParams))
	for i, tp := range toolParams {
		tools[i] = anthropic.ToolUnionParam{OfTool: &tp}
	}

	return tools
}