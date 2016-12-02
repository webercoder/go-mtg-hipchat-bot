package lib

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

// HipChatResponse .
type HipChatResponse struct {
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

func (r HipChatResponse) createMessage(cards []DeckbrewServiceResponseItem) string {
	if len(cards) == 0 {
		return "No cards were found."
	}

	tm := &TemplateManager{}
	cardsHTML := make([]string, 0, len(cards))

	for _, card := range cards {
		var tempBuffer bytes.Buffer
		templateObject := struct {
			Name     string
			Cost     string
			TypeLine string
			Text     template.HTML
			Editions []DeckbrewServiceResponseItemEdition
		}{
			card.Name,
			card.Cost,
			generateTypeLine(card),
			template.HTML(card.Text),
			card.Editions,
		}
		err := tm.Execute("card.html", templateObject, &tempBuffer)
		if err != nil {
			continue
		}
		cardsHTML = append(cardsHTML, tempBuffer.String())
	}
	return strings.Join(cardsHTML, "<br><br>")
}

// NewHipChatResponse .
func NewHipChatResponse(resultSets map[string][]DeckbrewServiceResponseItem) *HipChatResponse {
	resp := &HipChatResponse{
		Color:         "gray",
		Notify:        false,
		MessageFormat: "html",
	}

	messages := make(map[string]string)
	for name, cards := range resultSets {
		messages[name] = resp.createMessage(cards)
	}

	if len(messages) == 1 {
		for _, msg := range messages {
			resp.Message = msg
			return resp
		}
	}

	finalMessages := make([]string, len(messages))
	i := 0
	for name, val := range messages {
		finalMessages[i] = fmt.Sprintf("<strong>=== Matches for \"%s\" ===</strong><br>%s", name, val)
		i = i + 1
	}
	resp.Message = strings.Join(finalMessages, "<br><br>")

	return resp
}
