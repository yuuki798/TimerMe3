package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/ginx/dbx"
	"github.com/yuuki798/TimerMe3/core/ginx/rdsx"
	"github.com/yuuki798/TimerMe3/core/middleware/cors"
	"github.com/yuuki798/TimerMe3/core/miniox"
	"github.com/yuuki798/TimerMe3/internal/router"
)

func GinInit() *gin.Engine {
	r := gin.Default()
	dbx.InitDB()
	rdsx.InitCache()

	r.Use(cors.Middleware())
	router.GenerateRouters(r)

	miniox.MinioInit()

	return r
}
