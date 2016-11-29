package lib

// Response .
type Response struct {
	Color         string `json:"color"`
	Message       string `json:"message"`
	Notify        bool   `json:"notify"`
	MessageFormat string `json:"message_format"`
}

// NewResponse .
func NewResponse(cards []MTGRetrieverResponseItem) *Response {
	return &Response{}
}

// {
//     "color": "green",
//     "message": "It's going to be sunny tomorrow! (yey)",
//     "notify": false,
//     "message_format": "text"
// }
