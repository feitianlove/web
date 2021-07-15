package web

import (
	"github.com/feitianlove/web/config"
	"github.com/feitianlove/web/store"
	"github.com/patrickmn/go-cache"
)

type ClientWeb struct {
	conf  *config.Config
	Store *store.Store
	Cache *cache.Cache // 内存临时缓存 (重启丢失，支持TTL)
}

func NewWebClient(conf *config.Config) (*ClientWeb, error) {
	db, err := store.NewStore(conf)
	if err != nil {
		return nil, err
	}
	cliWeb := &ClientWeb{Store: db}

	return cliWeb, nil
}
