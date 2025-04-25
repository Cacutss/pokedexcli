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
	commands := GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if _, ok := commands[input[0]]; ok {
			commands[input[0]].callback()
		} else {
			fmt.Println("Uknown command")
		}
	}
}
