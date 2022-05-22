package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"warp-demo/service"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

// GetUser /*根据用户id，用户昵称查询*/
func (u *User) GetUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Request.URL.Query().Get("id"), 10, 64)
	user := service.NewUserService()
	data, err := user.GetUserById(id)
	if data != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"status":  10000,
		"data":    data,
		"message": "查询成功",
	})
}
