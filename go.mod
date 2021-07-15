module github.com/feitianlove/web

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/casbin/casbin/v2 v2.31.10
	github.com/casbin/gorm-adapter/v3 v3.3.2
	github.com/feitianlove/golib v0.0.0-20210715152206-b5bed50cfe70
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.7.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
