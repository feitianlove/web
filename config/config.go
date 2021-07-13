package config

import (
	"github.com/BurntSushi/toml"
	"github.com/feitianlove/golib/common/logger"
)

type Config struct {
	CasBin       *CasBinConfig
	WebLog       *logger.LogConf
	WebAccessLog *logger.LogConf
	MysqlLog     *logger.LogConf
	CtrlLog      *logger.LogConf
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
		CasBin: &CasBinConfig{
			Username: "",
			Passwd:   "",
			Port:     0,
			Database: "",
		},
		WebLog: &logger.LogConf{
			LogLevel:      "",
			LogPath:       "",
			LogReserveDay: 0,
			ReportCaller:  false,
		},
		WebAccessLog: &logger.LogConf{
			LogLevel:      "",
			LogPath:       "",
			LogReserveDay: 0,
			ReportCaller:  false,
		},
		MysqlLog: &logger.LogConf{
			LogLevel:      "",
			LogPath:       "",
			LogReserveDay: 0,
			ReportCaller:  false,
		},
		CtrlLog: &logger.LogConf{
			LogLevel:      "",
			LogPath:       "",
			LogReserveDay: 0,
			ReportCaller:  false,
		},
	}
}
