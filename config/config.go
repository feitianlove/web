package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	CasBin *CasBinConfig
}

type CasBinConfig struct {
	Username string
	Passwd   string
	Host     string
	Port     int64
	Database string
}

func InitConfig() (*Config, error) {
	var config = defaultConfig()
	_, err := toml.DecodeFile("./etc/web.conf", config)
	if err != nil {
		return nil, err
	}
	return config, err
}
func defaultConfig() *Config {
	return &Config{
		&CasBinConfig{
			Username: "",
			Passwd:   "",
			Port:     0,
			Database: "",
		},
	}
}
