package main

import (
	"singo/conf"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	// 装载路由
	r := server.NewRouter()
	r.Run(":80")

	//r.RunTLS(":80", "./7914142_qxxa.top.pem", "./7914142_qxxa.top.key") //将r.run(":8000")修改即可
	//r.RunTLS(":80", "./server.crt", "./server.key") //将r.run(":8000")修改即可
	//7914142_qxxa.top_public.crt
}

//func main() {
//	GinHttps(true) // 这里false 表示 http 服务，非 https
//}
//
//func GinHttps(isHttps bool) error {
//
//	r := gin.Default()
//	r.GET("/test", func(c *gin.Context) {
//		c.String(200, "test for 【%s】", "https")
//	})
//
//	if isHttps {
//		r.Use(TlsHandler(8000))
//		//7914142_qxxa.top.pem
//		//7914142_qxxa.top_public.pem
//		return r.RunTLS(":"+strconv.Itoa(8000), "./7914142_qxxa.top.pem", "./7914142_qxxa.top.key")
//	}
//
//	return r.Run(":" + strconv.Itoa(8000))
//}
//
//func TlsHandler(port int) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		secureMiddleware := secure.New(secure.Options{
//			SSLRedirect: true,
//			SSLHost:     ":" + strconv.Itoa(port),
//		})
//		err := secureMiddleware.Process(c.Writer, c.Request)
//
//		// If there was an error, do not continue.
//		if err != nil {
//			return
//		}
//
//		c.Next()
//	}
//}
