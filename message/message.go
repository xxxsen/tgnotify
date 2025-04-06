package message

type IMessage interface {
	Kind() string
	Marshal() (string, error)
}
