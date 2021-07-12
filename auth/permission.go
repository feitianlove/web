package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/feitianlove/web/config"
	_ "github.com/go-sql-driver/mysql"
)

var CasBin *casbin.Enforcer

func Init(conf config.CasBinConfig) error {
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/", conf.Username, conf.Passwd, conf.Host, conf.Port)
	a, err := gormadapter.NewAdapter("mysql", mysql) // Your driver and data source.
	if err != nil {
		return err
	}
	e, err := casbin.NewEnforcer("./etc/casbin.conf", a)
	if err != nil {
		return err
	}
	CasBin = e
	err = CasBin.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

//
func CheckPolicy(role, url, method string) (bool, error) {
	return CasBin.Enforce(role, url, method)
}

func AddPolicy(role, url, method string) (bool, error) {
	return CasBin.AddPolicy(role, url, method)
}
