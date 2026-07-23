package main

import (
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var systemTxt = `
You are a transplier of a programming language called Enlang
It is basicly a programming language where you write code in English and it complies to machine code
Your job is to turn that "code" into Go (golang version 1.24)
Respond with that code in the first response block as text
`

func getClient() anthropic.Client {
	return anthropic.NewClient(
		option.WithAPIKey(os.Getenv("CLAUDE")),
	)
}