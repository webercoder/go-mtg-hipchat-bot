package lib

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

type (
	// HipChatResponse .
	HipChatResponse struct {
		Color         string `json:"color"`
		Message       string `json:"message"`
		Notify        bool   `json:"notify"`
		MessageFormat string `json:"message_format"`
	}

	// CardTemplateData .
	CardTemplateData struct {
		Name     string
		Cost     template.HTML
		TypeLine string
		Text     template.HTML
		Editions []DeckbrewServiceResponseItemEdition
	}
)

// GenerateTypeLine .
func (hcr HipChatResponse) generateTypeLine(card *DeckbrewServiceResponseItem) string {
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

func (hcr HipChatResponse) createCardTemplateData(card *DeckbrewServiceResponseItem) *CardTemplateData {
	// Replace newlines in Text with <br>
	card.Text = strings.Replace(card.Text, "\n", "<br>", -1)

	// Replace 2/W, 2/U, 2/B, 2/R, 2/G with 2W, 2U, 2B, 2R, 2G
	dualRegex, _ := regexp.Compile(`{(\w)/(\w)}`)
	card.Cost = dualRegex.ReplaceAllString(card.Cost, "{${1}${2}}")
	card.Text = dualRegex.ReplaceAllString(card.Text, "{${1}${2}}")

	// Replace icons with images in cost and text
	iconRegex, _ := regexp.Compile(`{([^\}]+)}`)
	card.Cost = iconRegex.ReplaceAllString(card.Cost, "<img alt=\"${1}\" src=\"http://pub.webercoder.com/mtg/${1}.png\">")
	card.Text = iconRegex.ReplaceAllString(card.Text, "<img alt=\"${1}\" src=\"http://pub.webercoder.com/mtg/${1}.png\">")

	return &CardTemplateData{
		card.Name,
		template.HTML(card.Cost),
		hcr.generateTypeLine(card),
		template.HTML(card.Text),
		card.Editions,
	}
}

func (hcr HipChatResponse) createMessage(cards []DeckbrewServiceResponseItem) string {
	if len(cards) == 0 {
		return "No cards were found."
	}

	tm := &TemplateManager{}
	cardsHTML := make([]string, 0, len(cards))

	for _, card := range cards {
		var tempBuffer bytes.Buffer
		err := tm.Execute("card.html", hcr.createCardTemplateData(&card), &tempBuffer)
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
