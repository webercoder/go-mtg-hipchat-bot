package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type (
	// DeckbrewService .
	DeckbrewService struct {
		URL string
	}

	// DeckbrewServiceResponseItem .
	DeckbrewServiceResponseItem struct {
		Name       string                               `json:"name"`
		ID         string                               `json:"id"`
		URL        string                               `json:"url"`
		StoreURL   string                               `json:"store_url"`
		Supertypes []string                             `json:"supertypes"`
		Types      []string                             `json:"types"`
		Subtypes   []string                             `json:"subtypes"`
		Colors     []string                             `json:"colors"`
		CMC        int                                  `json:"cmc"`
		Formats    map[string]string                    `json:"formats"`
		Cost       string                               `json:"cost"`
		Text       string                               `json:"text"`
		Editions   []DeckbrewServiceResponseItemEdition `json:"editions"`
	}

	// DeckbrewServiceResponseItemEdition .
	DeckbrewServiceResponseItemEdition struct {
		Set          string          `json:"set"`
		SetID        string          `json:"set_id"`
		Rarity       string          `json:"rarity"`
		Artist       string          `json:"artist"`
		MultiverseID int64           `json:"multiverse_id"`
		Flavor       string          `json:"flavor"`
		Number       string          `json:"number"`
		Layout       string          `json:"layout"`
		Price        json.RawMessage `json:"price"`
		URL          string          `json:"url"`
		ImageURL     string          `json:"image_url"`
		SetURL       string          `json:"set_url"`
		StoreURL     string          `json:"store_url"`
		HTMLURL      string          `json:"html_url"`
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

func (mtgr DeckbrewService) cleanResponse(resp []DeckbrewServiceResponseItem) {
	for i := range resp {
		resp[i].Text = strings.Replace(resp[i].Text, "\n", "<br>", -1)
	}
}

// NewDeckbrewService .
func NewDeckbrewService() *DeckbrewService {
	return &DeckbrewService{URL: "https://api.deckbrew.com/mtg/cards"}
}

// GetCardsByName .
func (mtgr DeckbrewService) GetCardsByName(name string) ([]DeckbrewServiceResponseItem, error) {
	url := mtgr.URL + "?name=" + url.QueryEscape(name)
	resp := make([]DeckbrewServiceResponseItem, 0)
	err := getJSON(url, &resp)
	if err != nil {
		return nil, err
	}

	mtgr.cleanResponse(resp)

	return resp, nil
}
