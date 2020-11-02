package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englisBot struct{}
type hindiBot struct{}

func main() {
	eb := englisBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englisBot) getGreeting() string {
	return "Hello"
}

func (hindiBot) getGreeting() string {
	return "नमस्ते"
}
