// Package config config parser
package config

import (
	"encoding/json"
	"os"

	"github.com/xxxsen/common/logger"
)

type Channel struct {
	ChatID int64  `json:"chatid"`
	Token  string `json:"token"`
}

type Config struct {
	Listen    string            `json:"listen"`
	ChatID    int64             `json:"chatid"`
	Token     string            `json:"token"`
	User      map[string]string `json:"users"`
	LogConfig logger.LogConfig  `json:"log_config"`
}

func Parse(f string) (*Config, error) {
	raw, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	c := &Config{
		Listen: ":9902",
		LogConfig: logger.LogConfig{
			Level:   "debug",
			Console: true,
		},
	}
	if err := json.Unmarshal(raw, c); err != nil {
		return nil, err
	}
	return c, nil
}
