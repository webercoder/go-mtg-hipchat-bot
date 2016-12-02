package lib

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// MaxCards .
const MaxCards int = 5

// HipChatRequestController .
type HipChatRequestController struct{}

// GetNamesFromRequest .
func (rc HipChatRequestController) GetNamesFromRequest(r *HipChatRequest) []string {
	msg := r.Item.Message.Message
	var i int
	inlineRegexp, _ := regexp.Compile(`\[\[[^\]]+\]\]`)

	if inlineRegexp.MatchString(msg) { // Match [[card]] style first
		results := inlineRegexp.FindAllString(msg, -1)
		queries := make([]string, len(results))
		for i, val := range results {
			queries[i] = val[2 : len(val)-2]
		}
		return queries
	} else if i = strings.Index(msg, "/mtg"); i >= 0 { // Match /mtg card-name-here style next
		return []string{msg[i+len("/mtg ") : len(msg)]}
	}

	return []string{msg}
}

// HandleRequest .
func (rc HipChatRequestController) HandleRequest(r *http.Request) *HipChatResponse {
	hcreq, err := NewHipChatRequest(r)
	if err != nil {
		return &HipChatResponse{
			Color:         "red",
			Notify:        false,
			MessageFormat: "text",
			Message:       "Could not parse the HipChat request (damn you Atlassian!)",
		}
	}
	dbsvc := NewDeckbrewService()
	names := rc.GetNamesFromRequest(hcreq)
	fmt.Printf("Got: '%s'; Parsed: %+v\n", hcreq.Item.Message.Message, names)
	cards := dbsvc.GetCardsByNames(names, MaxCards)
	return NewHipChatResponse(cards)
}
