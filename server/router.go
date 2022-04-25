package server

import (
	"os"
	"singo/api/apiV1"
	"singo/api/apiV2"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", apiV1.Ping)

		// 用户登录
		v1.POST("user/register", apiV1.UserRegister)

		// 用户登录
		v1.POST("user/login", apiV1.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", apiV1.UserMe)
			auth.DELETE("user/logout", apiV1.UserLogout)
		}
	}
	v2 := r.Group("/api/v2")
	{
		v2.POST("pings", apiV2.Pings)

	}

	return r
}
