package main

import (
	"fmt"
	pokecache "github.com/Cacutss/pokedexcli/internal/pokecache"
	pokeapi "github.com/Cacutss/pokedexcli/pokeapi"
	"io"
	"os"
	"time"
)

type Config struct {
	Next *string
	Prev *string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *Config
	cache       *pokecache.Cache
}

type MapStruct struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = pokecache.NewCache(time.Minute * 1)
var Commands = make(map[string]cliCommand)

func CliMap() error {
	var url string
	if Commands["map"].config.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *Commands["map"].config.Next
	}
	var body []byte
	body, ok := Commands["mapb"].cache.Get(url)
	if !ok {
		res, err := pokeapi.GetRes(url)
		if err != nil {
			return fmt.Errorf("error getting response:%w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body")
		}
	}
	Map := MapStruct{}
	if err := pokeapi.UnmarshalBody(body, &Map); err != nil {
		return err
	}
	for _, v := range Map.Results {
		fmt.Println(v.Name)
	}
	if Map.Next != "" {
		nextURL := Map.Next
		Commands["map"].config.Next = &nextURL
	} else {
		Commands["map"].config.Next = nil
	}
	if Map.Previous != nil {
		prevURL := *Map.Previous
		Commands["map"].config.Prev = &prevURL
	} else {
		Commands["map"].config.Prev = nil
	}
	return nil
}

func CliMapb() error {
	var url string
	if Commands["mapb"].config.Prev == nil {
		fmt.Println("You're on the first page.")
		return nil
	} else {
		url = *Commands["mapb"].config.Prev
	}
	var body []byte
	body, ok := Commands["mapb"].cache.Get(url)
	if !ok {
		res, err := pokeapi.GetRes(url)
		if err != nil {
			return fmt.Errorf("error getting response:%w", err)
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body")
		}
	}
	Map := MapStruct{}
	if err := pokeapi.UnmarshalBody(body, &Map); err != nil {
		return err
	}
	for _, v := range Map.Results {
		fmt.Println(v.Name)
	}
	if Map.Next != "" {
		nextURL := &Map.Next
		Commands["mapb"].config.Next = nextURL
	} else {
		Commands["mapb"].config.Next = nil
	}
	if Map.Previous != nil {
		previousURL := Map.Previous
		Commands["mapb"].config.Prev = previousURL
	} else {
		Commands["mapb"].config.Prev = nil
	}
	return nil
}

func CliHelp() error {
	message := "Welcome to the Pokedex!\nUsage:\n\n"
	for k, v := range Commands {
		message += k + ": " + v.description + "\n"
	}
	fmt.Println(message)
	return nil
}

func CliExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func init() {
	Commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a message",
		callback:    CliHelp,
		config:      nil,
		cache:       cache,
	}
	Commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    CliExit,
		config:      nil,
		cache:       cache,
	}
	Commands["map"] = cliCommand{
		name:        "map",
		description: "Shows you the next 20 locations",
		callback:    CliMap,
		config:      &Config{},
		cache:       cache,
	}
	Commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Shows you the previous 20 locations",
		callback:    CliMapb,
		config:      Commands["map"].config,
		cache:       cache,
	}
}
