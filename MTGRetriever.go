package main

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
func (mtgr MTGRetriever) GetCard(name string) *MTGRetrieverResponse {
	return &MTGRetrieverResponse{}
}
