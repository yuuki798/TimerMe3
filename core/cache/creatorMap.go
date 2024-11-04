package cache

import (
	"github.com/trancecho/mundo-be-template/config"
	"github.com/trancecho/mundo-be-template/core/cache/driver"
	"github.com/trancecho/mundo-be-template/core/cache/types"
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
