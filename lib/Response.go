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
			Name       string
			Cost       string
			Supertypes []string
			Types      []string
			Subtypes   []string
			Text       template.HTML
		}{
			card.Name,
			card.Cost,
			card.Supertypes,
			card.Types,
			card.Subtypes,
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
