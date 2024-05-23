package main

import (
	"fmt"
	"strings"
)

const languageChosenENG = "Hello, "
const languageChosenPTBR = "Olá, "
const languageChosenESP = "Hola, "

func Hello(language, name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(ChoseLeanguage(language)+"%s", name)
}

func ChoseLeanguage(language string) (languageChosen string) {
	ESP := "ESP"
	PTBR := "PTBR"

	switch language {
	case ESP:
		languageChosen = languageChosenESP
	case PTBR:
		languageChosen = languageChosenPTBR
	default:
		languageChosen = languageChosenENG
	}
	return
}

func main() {
	fmt.Println(Hello("PTBR", " "))
}
