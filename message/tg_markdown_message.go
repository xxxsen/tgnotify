package message

type tgMarkdownMessage struct {
	text string
}

func (t *tgMarkdownMessage) Kind() string {
	return MKindTGMarkdown
}

func (t *tgMarkdownMessage) Marshal() (string, error) {
	return t.text, nil
}

func NewTGMarkdownMessage(text string) IMessage {
	return &tgMarkdownMessage{text: text}
}
