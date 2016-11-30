package lib

import "fmt"

// Response .
type Response struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        bool   `json:"notify"`
	MessageFormat string `json:"message_format"`
}

// NewResponse .
func NewResponse(cards []DeckbrewServiceResponseItem) *Response {
	resp := &Response{
		Color:         "green",
		Notify:        false,
		MessageFormat: "html",
	}

	str := fmt.Sprintf("Results (%d):<br>", len(cards))
	for _, card := range cards {
		str += card.Name + "<br>"
	}
	resp.Message = str

	return resp
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
