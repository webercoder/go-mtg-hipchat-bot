package lib

import (
	"bytes"
	"html/template"
	"strings"
)

// Response .
type Response struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        bool   `json:"notify"`
	MessageFormat string `json:"message_format"`
}

// GenerateTypeLine .
func generateTypeLine(card DeckbrewServiceResponseItem) string {
	parts := make([]string, 0, 3)
	if len(card.Supertypes) > 0 {
		parts = append(parts, strings.Join(card.Supertypes, " "))
	}
	if len(card.Types) > 0 {
		parts = append(parts, strings.Join(card.Types, " "))
	}
	if len(card.Subtypes) > 0 {
		parts = append(parts, "- "+strings.Join(card.Subtypes, " "))
	}
	return strings.Title(strings.Join(parts, " "))
}

// NewResponse .
func NewResponse(cards []DeckbrewServiceResponseItem) (*Response, error) {
	resp := &Response{
		Color:         "gray",
		Notify:        false,
		MessageFormat: "html",
	}
	tm := &TemplateManager{}
	cardsHTML := make([]string, len(cards))

	for i, card := range cards {
		var tempBuffer bytes.Buffer
		templateObject := struct {
			Name     string
			Cost     string
			TypeLine string
			Text     template.HTML
		}{
			card.Name,
			card.Cost,
			generateTypeLine(card),
			template.HTML(card.Text),
		}
		err := tm.Execute("card.html", templateObject, &tempBuffer)
		if err != nil {
			return nil, err
		}
		cardsHTML[i] = tempBuffer.String()
	}
	resp.Message = strings.Join(cardsHTML, "<br><br>")

	return resp, nil
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
