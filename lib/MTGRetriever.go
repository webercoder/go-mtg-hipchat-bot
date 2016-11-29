package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// MTGRetriever .
	MTGRetriever struct {
		URL string
	}

	// MTGRetrieverResponseItem .
	MTGRetrieverResponseItem struct {
		Name     string            `json:"name"`
		ID       string            `json:"id"`
		URL      string            `json:"url"`
		StoreURL string            `json:"store_url"`
		Types    []string          `json:"types"`
		Colors   []string          `json:"colors"`
		CMC      int               `json:"cmc"`
		Formats  map[string]string `json:"formats"`
		Cost     string            `json:"cost"`
		Text     string            `json:"text"`
		Editions json.RawMessage   `json:"editions"`
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

// NewMTGRetriever .
func NewMTGRetriever(url string) *MTGRetriever {
	return &MTGRetriever{URL: url}
}

// GetCards .
func (mtgr MTGRetriever) GetCardsByName(name string) ([]MTGRetrieverResponseItem, error) {
	url := mtgr.URL + "?name=" + name
	resp := make([]MTGRetrieverResponseItem, 0)
	err := getJSON(url, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
