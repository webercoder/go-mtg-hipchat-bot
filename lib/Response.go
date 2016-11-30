package lib

import (
	"bytes"
	"fmt"
)

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
		Color:         "green",
		Notify:        false,
		MessageFormat: "html",
	}
	tm := &TemplateManager{}
	var cardsHTML bytes.Buffer

	for _, card := range cards {
		err := tm.Execute("card.html", card, &cardsHTML)
		if err != nil {
			return nil, err
		}
	}
	resp.Message = fmt.Sprintf("<strong>%s</strong><ol>%s</ol>", fmt.Sprintf("Top %d Results:<br>", len(cards)), cardsHTML.String())

	return resp, nil
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
