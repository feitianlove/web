package logger

import (
	goliblogger "github.com/feitianlove/golib/common/logger"
	"github.com/feitianlove/web/config"
	"github.com/sirupsen/logrus"
)

var WebLog *logrus.Logger
var WebAccessLog *logrus.Logger
var CtrlLog *logrus.Logger
var MysqlLog *logrus.Logger

func init() {
	WebLog = goliblogger.NewLoggerInstance()
	WebAccessLog = goliblogger.NewLoggerInstance()
	CtrlLog = goliblogger.NewLoggerInstance()
	MysqlLog = goliblogger.NewLoggerInstance()
}

func InitWebLog(conf *goliblogger.LogConf) error {
	log, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	WebLog = log
	return nil
}
func InitWebAccessLog(conf *goliblogger.LogConf) error {
	log, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	WebAccessLog = log
	return nil
}
func InitMysqlLog(conf *goliblogger.LogConf) error {
	log, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	MysqlLog = log
	return nil
}
func InitCtrlLog(conf *goliblogger.LogConf) error {
	log, err := goliblogger.InitLogger(conf)
	if err != nil {
		return err
	}
	CtrlLog = log
	return nil
}

func InitLog(conf *config.Config) error {
	err := InitCtrlLog(conf.CtrlLog)
	if err != nil {
		return err
	}
	err = InitWebAccessLog(conf.WebAccessLog)
	if err != nil {
		return err
	}
	err = InitMysqlLog(conf.MysqlLog)
	if err != nil {
		return err
	}
	err = InitWebLog(conf.WebLog)
	if err != nil {
		return err
	}
	return nil
}
