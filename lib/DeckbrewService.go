package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func (mtgr DeckbrewService) titleCaseStrArray(arr []string) []string {
	return strings.Split(strings.Title(strings.Join(arr, " ")), " ")
}

func (mtgr DeckbrewService) cleanResponse(resp []DeckbrewServiceResponseItem) {
	for i := range resp {
		resp[i].Subtypes = mtgr.titleCaseStrArray(resp[i].Subtypes[:])
		resp[i].Types = mtgr.titleCaseStrArray(resp[i].Types[:])
		resp[i].Supertypes = mtgr.titleCaseStrArray(resp[i].Supertypes[:])
	}
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

	mtgr.cleanResponse(resp)

	return resp, nil
}
