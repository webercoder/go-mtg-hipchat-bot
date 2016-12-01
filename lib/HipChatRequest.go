package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// HipChatRequest .
	HipChatRequest struct {
		Event     string             `json:"event"`
		Item      HipChatRequestItem `json:"item"`
		WebhookID int                `json:"webhook_id"`
	}

	// HipChatRequestItem .
	HipChatRequestItem struct {
		Message HipChatRequestMessage `json:"message"`
		Room    HipChatRequestRoom    `json:"room"`
	}

	// HipChatRequestMessage .
	HipChatRequestMessage struct {
		Date     string `json:"date"`
		From     User   `json:"from"`
		ID       string `json:"id"`
		Mentions []User `json:"mentions"`
		Message  string `json:"message"`
		Type     string `json:"type"`
	}

	// HipChatRequestRoom .
	HipChatRequestRoom struct {
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

// NewHipChatRequest .
func NewHipChatRequest(r *http.Request) (*HipChatRequest, error) {
	var req HipChatRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("Difficulty parsing HipChat request: %+v", err)
	}
	return &req, nil
}
