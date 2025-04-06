package server

type config struct {
	addr  string
	users map[string]string
}

type Option func(c *config)

func WithUser(m map[string]string) Option {
	return func(c *config) {
		c.users = m
	}
}
