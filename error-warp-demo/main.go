package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"warp-demo/basic"
	"warp-demo/routes"
)

func main() {
	/*加载环境变量*/
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	basic.Init()
	router := gin.Default()
	//加载路由
	routes.Init(router)
	endless.ListenAndServe("127.0.0.1:8080", router)
}
