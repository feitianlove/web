package store

import (
	"fmt"
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Store struct {
	db *gorm.DB
}

func NewStore(conf *config.Config) (*Store, error) {

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.MysqlConf.User,
			conf.MysqlConf.Passwd,
			conf.MysqlConf.Host,
			conf.MysqlConf.Port,
			conf.MysqlConf.Database,
		))
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(10)
	db.LogMode(true)
	db.SetLogger(logger.MysqlLog)
	db.DB().SetConnMaxLifetime(60 * time.Second)
	db.AutoMigrate()
	return &Store{db: db}, nil
}
