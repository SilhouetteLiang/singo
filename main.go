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
	//r.Run(":80")

	r.RunTLS(":80","./server.crt","./server.key")   //将r.run(":8000")修改即可
}
