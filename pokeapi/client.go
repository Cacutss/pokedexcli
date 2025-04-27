package pokeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	//"log"
	"net/http"
)

func GetRes(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error responding from server: %w", err)
	}
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}
	return res, nil
}

func CreateHash(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	realhash := hash.Sum(nil)
	return hex.EncodeToString(realhash) + ".json"
}

func GetCacheDir() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can't find home path")
	}
	fullPath := filepath.Join(path, ".cache/pokedexcli")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.MkdirAll(fullPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return fullPath, nil
}

func GetCache(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func StoreCache(url string, body []byte) error {
	filename := CreateHash(url)
	dir, err := GetCacheDir()
	if err != nil {
		return fmt.Errorf("error no directory provided: %w", err)
	}
	fullPath := filepath.Join(dir, filename)
	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		return nil
	}
	err = os.WriteFile(fullPath, body, 0755)
	if err != nil {
		return err
	}
	return err
}

func UnmarshalBody(body []byte, ptr any) error {
	if body == nil {
		return fmt.Errorf("error no body passed")
	}
	if err := json.Unmarshal(body, ptr); err != nil {
		return fmt.Errorf("error unmarshaling: %w", err)
	}
	return nil
}
