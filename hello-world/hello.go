package main

import (
	"fmt"
)

const (
	bulgarian = "Bulgarian"
	french    = "French"
	spanish   = "Spanish"

	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	bulgarianHelloPrexix = "Zdravei, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case bulgarian:
		prefix = bulgarianHelloPrexix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("string", "string"))
}
