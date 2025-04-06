package model

type SendMessageRequest struct {
	Message     string `json:"message" binding:"required"`
	MessageType string `json:"message_type"`
}
