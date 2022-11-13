package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const japaneseHelloPrefix = "Konnichiwa, "

const spanishLanguage = "Spanish"
const japaneseLanguage = "Japanese"

const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case japaneseLanguage:
		prefix = japaneseHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
