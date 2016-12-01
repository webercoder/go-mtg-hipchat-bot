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
func NewHipChatResponse(resultSets [][]DeckbrewServiceResponseItem) *HipChatResponse {
	resp := &HipChatResponse{
		Color:         "gray",
		Notify:        false,
		MessageFormat: "html",
	}

	messages := make([]string, len(resultSets))
	for i, cards := range resultSets {
		messages[i] = resp.createMessage(cards)
	}

	if len(messages) == 1 {
		resp.Message = messages[0]
		return resp
	}

	for i, val := range messages {
		messages[i] = fmt.Sprintf("<strong>=== Matched Cards for Query #%d ===</strong><br>%s", i+1, val)
	}
	resp.Message = strings.Join(messages, "<br><br>")

	return resp
}
