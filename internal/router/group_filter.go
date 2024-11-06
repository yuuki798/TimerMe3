package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuki798/TimerMe3/core/builder"
	"github.com/yuuki798/TimerMe3/core/middleware/response"
	"github.com/yuuki798/TimerMe3/core/middleware/web"
	"github.com/yuuki798/TimerMe3/internal/router/protected"
)

func GenerateRouters(r *gin.Engine) *gin.Engine {

	newGroup := &builder.MyGroup{
		G: r.Group("/"),
	}

	baseGroup := builder.NewGroupBuilder().
		SetName("base").
		AddRoute("").
		SetFatherGroup(newGroup).
		AddMiddleware(response.ResponseMiddleware()).
		SetRoutes(Entity{}.Router).
		Build()

	builder.GetMyGroupDetail(baseGroup)
	{
		// 继承/base
		protectedGroup := builder.NewGroupBuilder().
			SetName("protected").
			SetFatherGroup(baseGroup).
			AddRoute("/api").
			AddMiddleware(web.JWTAuthMiddleware()). // 使用 JWTAuthMiddleware 中间件
			SetRoutes(protected.Entity{}.Router).
			Build()

		builder.GetMyGroupDetail(protectedGroup)
	}
	return r
}
