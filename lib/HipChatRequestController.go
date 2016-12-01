package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// HipChatRequestController .
type HipChatRequestController struct {
}

func (rc HipChatRequestController) constructNameQuery(str string) string {
	return fmt.Sprintf("?name=%s", url.QueryEscape(str))
}

// GetQueryFromRequest .
func (rc HipChatRequestController) GetQueryFromRequest(r *HipChatRequest) []string {
	msg := r.Item.Message.Message
	var i int
	inlineRegexp, _ := regexp.Compile(`\[\[[^\]]+\]\]`)

	if inlineRegexp.MatchString(msg) {
		results := inlineRegexp.FindAllString(msg, -1)
		queries := make([]string, len(results))
		for i, val := range results {
			queries[i] = rc.constructNameQuery(val[2 : len(val)-2])
		}
		return queries
	} else if i = strings.Index(msg, "/mtg"); i >= 0 {
		return []string{rc.constructNameQuery(msg[i+len("/mtg ") : len(msg)])}
	}
	return []string{rc.constructNameQuery(msg)}
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
	mtgr := NewDeckbrewService()
	queries := rc.GetQueryFromRequest(hcreq)
	fmt.Printf("Got: '%s'; Parsed: %+v\n", hcreq.Item.Message.Message, queries)
	cards := mtgr.GetCardsByQueries(queries)
	return NewHipChatResponse(cards)
}
