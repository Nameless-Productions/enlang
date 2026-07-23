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
Respond with that code in the first response block as text, you do not need any text like "i'm going to make this" as the user won't see anything except the code output in the FIRST block, use tools in the second and after
Just split out the code, do NOT use a codeblock
`

func getClient() anthropic.Client {
	return anthropic.NewClient(
		option.WithAPIKey(os.Getenv("CLAUDE")),
	)
}