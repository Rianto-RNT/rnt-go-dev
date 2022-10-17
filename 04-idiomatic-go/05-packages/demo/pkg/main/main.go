package main

import (
	"rnt-go-development/04-idiomatic-go/05-packages/demo/pkg/display"
	"rnt-go-development/04-idiomatic-go/05-packages/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from RNT Display")
	msg.Exciting("an Exciting message")
}
