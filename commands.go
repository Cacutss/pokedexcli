package main

import (
	"fmt"
	pokecache "github.com/Cacutss/pokedexcli/internal/pokecache"
	pokeapi "github.com/Cacutss/pokedexcli/pokeapi"
	"io"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	Next *string
	Prev *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
	Config      *Config
	Cache       *pokecache.Cache
	Params      []string
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

type PokemonInfo struct {
	Pokemon Pokemon `json:"pokemon"`
}

type MapZoneInfo struct {
	Encounters []PokemonInfo `json:"pokemon_encounters"`
}

var cache = pokecache.NewCache(time.Minute * 5)
var Commands = make(map[string]*cliCommand)

func CliHelp() error {
	message := "Welcome to the Pokedex!\nUsage:\nCommands that accept parameters will accept multiple parameters separated by a space And will stop execution upon first invalid parameter.\n\n"
	for k, v := range Commands {
		message += k + ": " + v.Description + "\n"
	}
	fmt.Println(message)
	return nil
}

func CliExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	SaveFile()
	os.Exit(0)
	return nil
}

func CliMap() error {
	var url string
	if Commands["map"].Config.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *Commands["map"].Config.Next
	}
	var body []byte
	body, ok := Commands["mapb"].Cache.Get(url)
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
		Commands["map"].Config.Next = &nextURL
	} else {
		Commands["map"].Config.Next = nil
	}
	if Map.Previous != nil {
		prevURL := *Map.Previous
		Commands["map"].Config.Prev = &prevURL
	} else {
		Commands["map"].Config.Prev = nil
	}
	return nil
}

func CliMapb() error {
	var url string
	if Commands["mapb"].Config.Prev == nil {
		fmt.Println("You're on the first page.")
		return nil
	} else {
		url = *Commands["mapb"].Config.Prev
	}
	var body []byte
	body, ok := Commands["mapb"].Cache.Get(url)
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
		Commands["mapb"].Config.Next = nextURL
	} else {
		Commands["mapb"].Config.Next = nil
	}
	if Map.Previous != nil {
		previousURL := Map.Previous
		Commands["mapb"].Config.Prev = previousURL
	} else {
		Commands["mapb"].Config.Prev = nil
	}
	return nil
}

func CliExplore() error {
	var url string
	if len(Commands["explore"].Params) < 1 {
		fmt.Println("not enough parameters")
		return nil
	}
	for _, v := range Commands["explore"].Params {
		url = "https://pokeapi.co/api/v2/location-area/" + v
		body, ok := Commands["explore"].Cache.Get(url)
		if !ok {
			res, err := pokeapi.GetRes(url)
			if err != nil {
				return fmt.Errorf("error fetching response: %w", err)
			}
			defer res.Body.Close()
			if res.StatusCode == 404 {
				fmt.Println("Not found")
				return nil
			}
			body, err = io.ReadAll(res.Body)
			Commands["explore"].Cache.Add(url, body)
		}
		fmt.Printf("Exploring %s...\n", v)
		zoneInfo := MapZoneInfo{}
		if err := pokeapi.UnmarshalBody(body, &zoneInfo); err != nil {
			return fmt.Errorf("error unmarshaling body: %w", err)
		}
		fmt.Println("Found Pokemon:")
		for _, v := range zoneInfo.Encounters {
			fmt.Println(v.Pokemon.Name)
		}
	}
	return nil
}

func CliCatch() error {
	if len(Commands["catch"].Params) < 1 {
		fmt.Println("not enough parameters")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + Commands["catch"].Params[0]
	cache := Commands["catch"].Cache
	var body []byte
	body, ok := cache.Get(url)
	if !ok {
		res, err := pokeapi.GetRes(url)
		if err != nil {
			return fmt.Errorf("error fetching response: %w", err)
		}
		if res.StatusCode == 404 {
			return fmt.Errorf("error pokemon doesn't exist")
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error reading body response: %w", err)
		}
	}
	var pokemon Pokemon
	if err := pokeapi.UnmarshalBody(body, &pokemon); err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	for i := 1; i < 4; i++ {
		random := rand.Intn(600)
		if random > pokemon.BaseExp {
			fmt.Printf("%d", i)
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
			time.Sleep(time.Millisecond * 250)
			fmt.Println(".")
		} else {
			fmt.Printf("%s", pokemon.Name)
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
			time.Sleep(time.Millisecond * 250)
			fmt.Print(".")
			fmt.Println("  Broke free!")
			return nil
		}
	}
	fmt.Printf("%s", pokemon.Name)
	time.Sleep(time.Millisecond * 500)
	fmt.Print(".")
	time.Sleep(time.Millisecond * 500)
	fmt.Print(".")
	time.Sleep(time.Millisecond * 500)
	fmt.Print(".")
	fmt.Println("  Succesfully catched!")
	if _, ok := Pokedex[pokemon.Name]; !ok {
		Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s added to the pokedex!\n", pokemon.Name)
	}
	if err := SaveFile(); err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func init() {
	Commands["help"] = &cliCommand{
		Name:        "help",
		Description: "Displays a message with all commands",
		Callback:    CliHelp,
		Config:      nil,
		Cache:       nil,
		Params:      nil,
	}
	Commands["exit"] = &cliCommand{
		Name:        "exit",
		Description: "Exit the pokedex",
		Callback:    CliExit,
		Config:      nil,
		Cache:       nil,
		Params:      nil,
	}
	Commands["map"] = &cliCommand{
		Name:        "map",
		Description: "Shows you the next 20 locations",
		Callback:    CliMap,
		Config:      &Config{},
		Cache:       cache,
		Params:      nil,
	}
	Commands["mapb"] = &cliCommand{
		Name:        "mapb",
		Description: "Shows you the previous 20 locations",
		Callback:    CliMapb,
		Config:      Commands["map"].Config,
		Cache:       cache,
		Params:      nil,
	}
	Commands["explore"] = &cliCommand{
		Name:        "explore",
		Description: "Explore <area_name> shows you the pokemon in the area",
		Callback:    CliExplore,
		Config:      nil,
		Cache:       cache,
		Params:      make([]string, 0),
	}
	Commands["catch"] = &cliCommand{
		Name:        "catch",
		Description: "Throws a pokeball to catch <pokemon> and tries to catch it, on sucessful attempts adds the pokemon to your pokedex",
		Callback:    CliCatch,
		Config:      nil,
		Cache:       cache,
		Params:      make([]string, 0),
	}
}
