package main

import (
	"fmt"
)

func writeColor(str string, color string, newline bool) {
	colors := map[string]string{
		"reset":  "\033[0m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
	}

	selectedColor := colors[color]

	print(selectedColor)

	if newline {
		fmt.Println(string(colors[selectedColor]), str, string(colors["reset"]))
	} else {
		fmt.Print(string(colors[selectedColor]), str, string(colors["reset"]))
	}

}
