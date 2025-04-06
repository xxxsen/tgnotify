package message

type textMessage struct {
	text string
}

func (t *textMessage) Kind() string {
	return MKindText
}

func (t *textMessage) Marshal() (string, error) {
	return t.text, nil
}

func NewTextMessage(text string) IMessage {
	return &textMessage{
		text: text,
	}
}
