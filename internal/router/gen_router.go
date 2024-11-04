package router

import (
	"github.com/gin-gonic/gin"
	"github.com/trancecho/mundo-be-template/core/middleware/response"
	"github.com/trancecho/mundo-be-template/core/middleware/web"
	"github.com/trancecho/mundo-be-template/internal/router/protected"
)

func GenerateRouters(r *gin.Engine) *gin.Engine {

	newGroup := &MyGroup{
		g: r.Group("/"),
	}

	baseGroup := NewGroupBuilder().
		SetName("base").
		AddRoute("").
		SetFatherGroup(newGroup).
		AddMiddleware(response.ResponseMiddleware()).
		SetRoutes(Entity{}.Router).
		Build()

	GetMyGroupDetail(baseGroup)
	{
		// 继承/base
		protectedGroup := NewGroupBuilder().
			SetName("protected").
			SetFatherGroup(baseGroup).
			AddRoute("/api").
			AddMiddleware(web.JWTAuthMiddleware()). // 使用 JWTAuthMiddleware 中间件
			SetRoutes(protected.Entity{}.Router).
			Build()

		GetMyGroupDetail(protectedGroup)
	}
	return r
}
