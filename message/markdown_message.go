package message

import "fmt"

type markdownMessage struct {
	text string
}

func (m *markdownMessage) Kind() string {
	return MKindMarkdown
}

func (m *markdownMessage) Marshal() (string, error) {
	return "", fmt.Errorf("no impl")
}

func NewMarkdownMessage(text string) IMessage {
	return &markdownMessage{text: text}
}
