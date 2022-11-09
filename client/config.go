package client

type config struct {
	host string
	ak   string
	sk   string
}

type Option func(c *config)

func WithHost(h string) Option {
	return func(c *config) {
		c.host = h
	}
}

func WithUser(ak, sk string) Option {
	return func(c *config) {
		c.ak = ak
		c.sk = sk
	}
}
