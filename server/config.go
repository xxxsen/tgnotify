package server

type config struct {
	addr  string
	users map[string]string
}

type Option func(c *config)

func WithBind(addr string) Option {
	return func(c *config) {
		c.addr = addr
	}
}

func WithUser(m map[string]string) Option {
	return func(c *config) {
		c.users = m
	}
}
