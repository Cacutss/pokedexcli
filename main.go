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
		if _, ok := Commands[input[0]]; ok {
			Commands[input[0]].callback()
		} else {
			fmt.Println("Uknown command")
		}
	}
}
