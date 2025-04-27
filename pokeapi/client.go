package pokeapi

import (
	"encoding/json"
	"fmt"
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

func UnmarshalBody(body []byte, ptr any) error {
	if body == nil {
		return fmt.Errorf("error no body passed")
	}
	if err := json.Unmarshal(body, ptr); err != nil {
		return fmt.Errorf("error unmarshaling: %w", err)
	}
	return nil
}
