package pokecache

import (
	"log"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second * 5)
	if cache == nil {
		log.Fatal("error initializing cache")
	}
	entry := []byte{12, 15, 200}
	entry2 := []byte{}
	url := "github.com/Cacutss/pokedexcli/internal/pokecache"
	cache.Add(url, entry)
	time.Sleep(5 * time.Second)
	cache.Add(url+"/pokemon", entry2)
	if len(cache.Entries) > 1 {
		log.Fatal("reaploop failure")
	}
}
