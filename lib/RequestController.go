package lib

import "strings"

// RequestController .
type RequestController struct {
}

// GetQueryFromRequest .
func (rc RequestController) GetQueryFromRequest(r *Request) string {
	msg := r.Item.Message.Message
	i := strings.Index(msg, "/mtg")
	if i >= 0 {
		return msg[i+5 : len(msg)]
	}
	return msg
}

// HandleRequest .
func (rc RequestController) HandleRequest(r *Request) (*Response, error) {
	mtgr := NewDeckbrewService()
	query := rc.GetQueryFromRequest(r)
	cards, err := mtgr.GetCardsByName(query)
	if err != nil {
		return nil, err
	}
	return NewResponse(cards)
}
