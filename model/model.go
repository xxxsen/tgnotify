package model

type SendMessageRequest struct {
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}
