package message

import "fmt"

type htmlMessage struct {
	text string
}

func (h *htmlMessage) Kind() string {
	return MKindHTML
}

func (h *htmlMessage) Marshal() (string, error) {
	return "", fmt.Errorf("no impl")
}

func NewHTMLMessage(text string) IMessage {
	return &htmlMessage{text: text}
}
