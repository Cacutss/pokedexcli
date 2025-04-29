package main

import (
	"bufio"
	"fmt"
	//"io"
	//"log"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Try again")
			continue
		}
		if _, ok := Commands[input[0]]; ok {
			if Commands[input[0]].Params != nil && len(input) > 1 {
				Commands[input[0]].Params = input[1:]
			}
			Commands[input[0]].Callback()
		} else {
			fmt.Println("Uknown command")
		}
	}
}
