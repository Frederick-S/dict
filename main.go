package main

import (
	"fmt"
	"os"
)

func display(word Word) {
	if len(word.explanations) == 0 {
		if len(word.suggestions) == 0 {
			fmt.Println("No results found")

			os.Exit(0)
		} else {
			fmt.Println("Are you looking for:")

			for _, suggestion := range word.suggestions {
				fmt.Println(suggestion)
			}
		}
	}

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

	display(word)
}
