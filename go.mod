module github.com/feitianlove/web

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/casbin/casbin/v2 v2.31.10
	github.com/casbin/gorm-adapter/v3 v3.3.2
	github.com/feitianlove/golib v0.0.0-20210415083646-83549b24dc6f
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/sirupsen/logrus v1.7.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
