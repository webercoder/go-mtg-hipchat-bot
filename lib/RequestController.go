package lib

// RequestController .
type RequestController struct {
}

// HandleRequest .
func (rc RequestController) HandleRequest(r *Request) *Response {
	mtgr := NewMTGRetriever("http://api.deckbrew.com/mtg/cards?name=")
	cards, err := mtgr.GetCardsByName(r.Item.Message.Message)
	if err != nil {
		return nil
	}
	return NewResponse(cards)
}
