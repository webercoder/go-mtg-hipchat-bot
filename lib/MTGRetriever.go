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

	// MTGRetrieverResponse .
	MTGRetrieverResponse struct {
		Name     string   `json:"name"`
		StoreURL string   `json:"store_url"`
		Types    []string `json:"types"`
		Cost     string   `json:"cost"`
		Text     string   `json:"text"`
	}
)

// NewMTGRetriever .
func NewMTGRetriever(url string) *MTGRetriever {
	return &MTGRetriever{URL: url}
}

// GetCard .
func (mtgr MTGRetriever) GetCard(name string) (*MTGRetrieverResponse, error) {
	url := mtgr.URL + name
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to %v", url)
	}

	decoder := json.NewDecoder(r.Body)
	var resp MTGRetrieverResponse
	err = decoder.Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("Could not parse the response from %v", url)
	}

	return &resp, nil
}
