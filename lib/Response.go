package lib

import (
	"bytes"
	"fmt"
)

// MaxCards .
const MaxCards int = 5

// Response .
type Response struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        bool   `json:"notify"`
	MessageFormat string `json:"message_format"`
}

// NewResponse .
func NewResponse(cards []DeckbrewServiceResponseItem) (*Response, error) {
	resp := &Response{
		Color:         "gray",
		Notify:        false,
		MessageFormat: "html",
	}
	tm := &TemplateManager{}
	var cardsHTML bytes.Buffer

	subset := cards[:MaxCards]
	for _, card := range subset {
		err := tm.Execute("card.html", card, &cardsHTML)
		if err != nil {
			return nil, err
		}
	}
	resp.Message = fmt.Sprintf("<strong>Top Matches</strong><br><br>%s", cardsHTML.String())

	return resp, nil
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
