package server

type config struct {
	addr   string
	chatid int64
	token  string
	users  map[string]string
}

type Option func(c *config)

func WithBind(addr string) Option {
	return func(c *config) {
		c.addr = addr
	}
}

func WithBotConfig(chatid int64, token string) Option {
	return func(c *config) {
		c.chatid = chatid
		c.token = token
	}
}

func WithUser(m map[string]string) Option {
	return func(c *config) {
		c.users = m
	}
}
