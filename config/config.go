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
	MysqlConf    *MysqlConf
	Web          *Web
	Master       *Master
}

type MysqlConf struct {
	User     string
	Passwd   string
	Host     string
	Port     int64
	Database string
}
type CasBinConfig struct {
	Username string
	Passwd   string
	Host     string
	Port     int64
	Database string
}
type Web struct {
	ListenPort int64
	Domain     string
	StaticDir  string
}

type Master struct {
	ListenPort int64
	Domain     string
	Token      string
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
			LogLevel:      "info",
			LogPath:       "/Users/fenghui/goCode/web/log/web.log",
			LogReserveDay: 1,
			ReportCaller:  true,
		},
		WebAccessLog: &logger.LogConf{
			LogLevel:      "info",
			LogPath:       "/Users/fenghui/goCode/web/log/web_access.log",
			LogReserveDay: 1,
			ReportCaller:  true,
		},
		MysqlLog: &logger.LogConf{
			LogLevel:      "info",
			LogPath:       "/Users/fenghui/goCode/web/log/mysql.log",
			LogReserveDay: 1,
			ReportCaller:  true,
		},
		CtrlLog: &logger.LogConf{
			LogLevel:      "info",
			LogPath:       "/Users/fenghui/goCode/web/log/ctrl.log",
			LogReserveDay: 1,
			ReportCaller:  true,
		},
		MysqlConf: &MysqlConf{
			User:     "",
			Passwd:   "",
			Host:     "",
			Port:     0,
			Database: "",
		},
	}
}
