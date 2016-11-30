package lib

type (
	// Request .
	Request struct {
		Event     string      `json:"event"`
		Item      RequestItem `json:"item"`
		WebhookID int         `json:"webhook_id"`
	}

	// RequestItem .
	RequestItem struct {
		Message RequestMessage `json:"message"`
		Room    RequestRoom    `json:"room"`
	}

	// RequestMessage .
	RequestMessage struct {
		Date     string `json:"date"`
		From     User   `json:"from"`
		ID       string `json:"id"`
		Mentions []User `json:"mentions"`
		Message  string `json:"message"`
		Type     string `json:"type"`
	}

	// RequestRoom .
	RequestRoom struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// User .
	User struct {
		ID          int    `json:"id"`
		MentionName string `json:"mention_name"`
		Name        string `json:"name"`
	}
)
