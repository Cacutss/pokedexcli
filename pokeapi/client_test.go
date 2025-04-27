package pokeapi

import (
	"fmt"
	"io"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://pokeapi.co/api/v2/pokemon/pikachu"
	res, err := GetRes(url)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("response is nil")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error reading body:%w", err)
	}
	if err := StoreCache(url, body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
}

func TestUnmarshal(t *testing.T) {
	var result string
	url := "https://pokeapi.co/api/v2/pokemon/drampa"
	res, err := GetRes(url)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("response is nil")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error reading body")
	}
	res.Body.Close()
	if err = UnmarshalBody(body, &result); err != nil {
		fmt.Println("this failed")
	}
	fmt.Println(result)
}
