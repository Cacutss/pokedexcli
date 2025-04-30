package main

import (
	"log"
	"testing"
)

func TestPokedexCache(t *testing.T) {
	if err := SaveFile(); err != nil {
		log.Fatal("error registering pokemon")
	}
}
