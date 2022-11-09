//Package config config parser
package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/logger"
)

type Config struct {
	Listen string            `json:"listen"`
	ChatID int64             `json:"chatid"`
	Token  string            `json:"token"`
	User   map[string]string `json:"users"`
	Log    logger.LogConfig  `json:"log"`
}

func Parse(f string) (*Config, error) {
	raw, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, errs.Wrap(errs.ErrIO, "read file fail", err)
	}
	c := &Config{}
	if err := json.Unmarshal(raw, c); err != nil {
		return nil, errs.Wrap(errs.ErrUnknown, "decode fail", err)
	}
	return c, nil
}
