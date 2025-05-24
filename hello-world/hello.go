package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"

	englishPrefix = "Hello "
	spanishPrefix = "Hola "
	frenchPrefix   = "Bonjour "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingLang(lang) + name
}

func greetingLang(lang string) (prefix string) {
	prefix = englishPrefix
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}
	return 
}

func main() {
	fmt.Println(Hello("Hamza", "Spanish"))
}
