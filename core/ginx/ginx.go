package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/database"
	"github.com/trancecho/mundo-be-template/core/middleware/cors"
	"github.com/trancecho/mundo-be-template/internal/app/ping"
	"github.com/trancecho/mundo-be-template/internal/router"
	"log"
)

func GinInit() *gin.Engine {
	r := gin.Default()
	db := database.GetDb("MainMysql")
	if db == nil {
		log.Fatalln("db not found")
	}
	err := db.AutoMigrate(
		&ping.TestModel{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	r.Use(cors.Middleware())
	router.GenerateRouters(r)

	return r
}
