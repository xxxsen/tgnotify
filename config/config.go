//Package config config parser
package config

import (
	"encoding/json"
	"io/ioutil"
)

//ServerConfig server config
type ServerConfig struct {
	Listen string `json:"listen"`
}

//DBConfig db config
type DBConfig struct {
	Host   string `json:"host"`
	Port   int16  `json:"port"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	DBName string `json:"dbname"`
}

//BotConfig bot config
type BotConfig struct {
	Token string `json:"token"`
}

//LogConfig log config
type LogConfig struct {
	File         string `json:"file"`
	Level        string `json:"level"`
	Count        int32  `json:"count"`
	Size         int32  `json:"size"`
	KeepDays     int32  `json:"keep_days"`
	WriteConsole bool   `json:"write_console"`
}

//NotifyConfig basic config
type NotifyConfig struct {
	Server ServerConfig `json:"server_config"`
	DB     DBConfig     `json:"db_config"`
	Bot    BotConfig    `json:"bot_config"`
	Log    LogConfig    `json:"log_config"`
}

var nc *NotifyConfig = &NotifyConfig{
	Log: LogConfig{
		File:         "./tgnotify.log",
		Level:        "debug",
		Count:        1,
		Size:         10485760,
		KeepDays:     1,
		WriteConsole: true,
	},
}

//Global get a global config instance
func Global() *NotifyConfig {
	return nc
}

//Parse parse config file
func (nc *NotifyConfig) Parse(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, nc)
	if err != nil {
		return err
	}
	return nil
}
