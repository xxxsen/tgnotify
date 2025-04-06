package message

type tgHTMLMessage struct {
	text string
}

func (t *tgHTMLMessage) Kind() string {
	return MKindTGRawHTML
}

func (t *tgHTMLMessage) Marshal() (string, error) {
	return t.text, nil
}

func NewTGHTMLMessage(text string) IMessage {
	return &tgHTMLMessage{text: text}
}
