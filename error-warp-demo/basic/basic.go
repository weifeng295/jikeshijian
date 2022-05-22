package basic

import (
	"warp-demo/basic/config"
	"warp-demo/basic/db"
	"warp-demo/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
