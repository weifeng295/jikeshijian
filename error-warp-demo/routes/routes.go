package routes

import (
	"github.com/gin-gonic/gin"
	"warp-demo/controllers"
)

type Routes interface {
}

func Init(router *gin.Engine) {
	user(router) //用户路由
}

/*用户信息*/
func user(router *gin.Engine) {
	user := controllers.NewUser()
	router.GET("user/get_user", user.GetUser)
}
