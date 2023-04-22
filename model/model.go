package model

type SendMessageRequest struct {
	Channel     string `json:"channel"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}
