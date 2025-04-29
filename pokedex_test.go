package main

import (
	"log"
	"testing"
)

func TestPokedexCache(t *testing.T) {
	magma := Pokemon{
		Name: "Magmasaurus",
		Url:  "example.com",
	}
	if err := RegisterPokemon(&magma); err != nil {
		log.Fatal("error registering pokemon")
	}
}
