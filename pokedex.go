package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	fullPath = "/.cache/pokedexcli"
	filePath = "/pokedex.json"
)

type PokemonApi struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Pokemon struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Stats  []struct {
		BaseStat int        `json:"base_stat"`
		Effort   int        `json:"effort"`
		Stat     PokemonApi `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int        `json:"slot"`
		Type PokemonApi `json:"type"`
	} `json:"types"`
	BaseExp int `json:"base_experience"`
}

var Pokedex = make(map[string]Pokemon)

func SaveFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home dir")
	}
	cachePath := homeDir + fullPath
	if _, err = os.Stat(cachePath); os.IsNotExist(err) {
		if err = os.MkdirAll(cachePath, 0755); err != nil {
			return err
		}
	}
	body, err := json.Marshal(Pokedex)
	if err != nil {
		return fmt.Errorf("error creating pokemon json: %w", err)
	}
	pathToFile := cachePath + filePath
	file, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("error creating the pokedex: %w", err)
	}
	defer file.Close()
	if _, err = file.Write(body); err != nil {
		return fmt.Errorf("error writing into json: %w", err)
	}
	return nil
}

func init() {
	homeDir, _ := os.UserHomeDir()
	if _, err := os.Stat(homeDir + fullPath); os.IsNotExist(err) {
		os.MkdirAll(homeDir+fullPath, 0755)
	}
	pathToFile := homeDir + fullPath + filePath
	var body []byte
	body, err := os.ReadFile(pathToFile)
	if os.IsNotExist(err) {
		file, _ := os.OpenFile(pathToFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
		defer file.Close()
		return
	}
	json.Unmarshal(body, &Pokedex)
}
