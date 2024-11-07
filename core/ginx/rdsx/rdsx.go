package rdsx

import (
	"github.com/yuuki798/TimerMe3/core/cache"
	"github.com/yuuki798/TimerMe3/core/cache/types"
	"log"
)

var Cache types.Cache

func InitCache() {
	Cache = cache.GetCache("MainRedis")
	if Cache == nil {
		log.Fatalln("fail to get cache")
	}
}
