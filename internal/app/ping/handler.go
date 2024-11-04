package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/cache"
	"github.com/trancecho/mundo-be-template/core/database"
	"github.com/trancecho/mundo-be-template/core/libx"
	"log"
	"time"
)

func Handler(c *gin.Context) {
	libx.Ok(c, "pong")
}

func TestMysql(c *gin.Context) {
	db := database.GetDb("MainMysql")
	if db == nil {
		libx.Err(c, 500, "db not found", nil)
		return
	}

	var testEntity TestModel
	// 如果不存在则创建
	err := db.FirstOrCreate(&testEntity, TestModel{
		Name: "test",
		Age:  18,
	}).Error

	if err != nil {
		libx.Err(c, 500, "db error", err)
		return
	}

	db.First(&testEntity)
	log.Println("got it! : ", testEntity)

	libx.Ok(c, "mysql pong！"+testEntity.Name+"Created At: "+testEntity.CreatedAt.String())
}

func TestRedis(c *gin.Context) {
	redis := cache.GetCache("MainRedis")
	if redis == nil {
		libx.Err(c, 500, "cache not found", nil)
		return
	}

	key := "test"
	value := "test"

	// 尝试从缓存中获取值
	var cachedValue string
	val, exists := redis.GetString(key)
	if exists == false {
		// 如果Key不存在或获取失败，就设置值
		if err := redis.Set(key, value, time.Second*60); err != nil {
			libx.Err(c, 500, "redis set error", err)
			return
		}
		cachedValue = value
	}

	if exists == true {
		cachedValue = val
	}

	libx.Ok(c, "redis pong！"+cachedValue)
}
