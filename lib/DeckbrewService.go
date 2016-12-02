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
func (dbsvc DeckbrewService) getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error retrieving %s: %+v", url, err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (dbsvc DeckbrewService) cleanResponse(resp []DeckbrewServiceResponseItem) {
	for i := range resp {
		resp[i].Text = strings.Replace(resp[i].Text, "\n", "<br>", -1)
	}
}

// NewDeckbrewService .
func NewDeckbrewService() *DeckbrewService {
	return &DeckbrewService{URL: "https://api.deckbrew.com/mtg/cards"}
}

func (dbsvc DeckbrewService) constructURL(queryMap map[string]string) string {
	u, _ := url.Parse(dbsvc.URL)
	q := u.Query()
	for i, val := range queryMap {
		q.Add(i, val)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (dbsvc DeckbrewService) getCardsByURL(url string, limit int) ([]DeckbrewServiceResponseItem, error) {
	resp := make([]DeckbrewServiceResponseItem, 0)
	err := dbsvc.getJSON(url, &resp)
	if err != nil {
		return nil, err
	}

	if limit > 0 && len(resp) > limit {
		resp = resp[0:limit]
	}

	dbsvc.cleanResponse(resp)

	return resp, nil
}

// GetCardsByNames .
func (dbsvc DeckbrewService) GetCardsByNames(names []string, limit int) map[string][]DeckbrewServiceResponseItem {
	results := make(map[string][]DeckbrewServiceResponseItem)
	for _, name := range names {
		url := dbsvc.constructURL(map[string]string{"name": name})
		resultSet, _ := dbsvc.getCardsByURL(url, limit)
		for j, item := range resultSet {
			if strings.ToLower(item.Name) == strings.ToLower(name) {
				resultSet = resultSet[j : j+1]
				break
			}
		}
		results[name] = resultSet
	}
	return results
}
