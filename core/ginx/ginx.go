package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/database"
	"github.com/yuuki798/TimerMe3/core/middleware/cors"
	"github.com/yuuki798/TimerMe3/core/miniox"
	"github.com/yuuki798/TimerMe3/internal/router"
	"log"
)

func GinInit() *gin.Engine {
	r := gin.Default()
	db := database.GetDb("MainMysql")
	if db == nil {
		log.Fatalln("db not found")
	}
	err := AutoMigrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	r.Use(cors.Middleware())
	router.GenerateRouters(r)

	miniox.MinioInit()

	return r
}
