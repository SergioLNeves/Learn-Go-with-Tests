package main

import (
	"fmt"
	"strings"
)

const sayHello = "Hello,"

func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s", sayHello, name)
}

func main() {
	fmt.Println(Hello(" "))
}
