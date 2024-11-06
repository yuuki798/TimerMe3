package cache

import (
	"github.com/yuuki798/TimerMe3/config"
	"github.com/yuuki798/TimerMe3/core/cache/driver"
	"github.com/yuuki798/TimerMe3/core/cache/types"
)

type Creator interface {
	Create(conf config.Cache) (types.Cache, error)
}

func init() {
	typeMap["redis"] = driver.RedisCreator{}
}

var typeMap = make(map[string]Creator)

func getCreatorByType(cacheType string) Creator {
	return typeMap[cacheType]
}
