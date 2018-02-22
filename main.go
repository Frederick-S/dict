package main

import (
	"fmt"
	"os"
)

func print(word Word) {
	for _, explanation := range word.explanations {
		fmt.Println(explanation)
	}
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Usage: dict word")

		os.Exit(0)
	}

	word := query(arguments[0])

	print(word)
}
