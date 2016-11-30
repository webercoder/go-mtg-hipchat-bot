package lib

import (
	"bytes"
	"html/template"
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
	var cardsHTML bytes.Buffer

	for _, card := range cards {
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
		err := tm.Execute("card.html", templateObject, &cardsHTML)
		if err != nil {
			return nil, err
		}
	}
	resp.Message = cardsHTML.String()

	return resp, nil
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
