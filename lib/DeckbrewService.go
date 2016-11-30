package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// DeckbrewService .
	DeckbrewService struct {
		URL string
	}

	// DeckbrewServiceResponseItem .
	DeckbrewServiceResponseItem struct {
		Name       string            `json:"name"`
		ID         string            `json:"id"`
		URL        string            `json:"url"`
		StoreURL   string            `json:"store_url"`
		Supertypes []string          `json:"supertypes"`
		Types      []string          `json:"types"`
		Subtypes   []string          `json:"subtypes"`
		Colors     []string          `json:"colors"`
		CMC        int               `json:"cmc"`
		Formats    map[string]string `json:"formats"`
		Cost       string            `json:"cost"`
		Text       string            `json:"text"`
		Editions   json.RawMessage   `json:"editions"`
	}
)

// http://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error retrieving %s: %+v", url, err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// NewDeckbrewService .
func NewDeckbrewService() *DeckbrewService {
	return &DeckbrewService{URL: "https://api.deckbrew.com/mtg/cards"}
}

// GetCardsByName .
func (mtgr DeckbrewService) GetCardsByName(name string) ([]DeckbrewServiceResponseItem, error) {
	url := mtgr.URL + "?name=" + name
	resp := make([]DeckbrewServiceResponseItem, 0)
	err := getJSON(url, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
