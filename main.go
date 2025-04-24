package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	lower := strings.ToLower(text)
	splitted := strings.Split(lower, " ")
	for _, word := range splitted {
		if len(word) != 0 {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	//scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
	}
}
