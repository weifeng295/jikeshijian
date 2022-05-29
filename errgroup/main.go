package main

import (
	"errgroup/service"
	"log"
	"net/http"
	"time"
)

//由于老师讲解的errgroup和context和chan知识点能理解，但是自己组合context、errgroup、chan写不出来，只能模仿别人好的资料写的。
func main() {
	s1 := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	s2 := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: nil,
	}
	s3 := &http.Server{
		Addr:    "127.0.0.1:8082",
		Handler: nil,
	}

	logger := log.Default()

	app := service.NewApp(
		service.WithServer(s1),
		service.WithServer(s2),
		service.WithServer(s3),
		service.WithLog(logger),
	)

	// 测试10s后退出。
	time.AfterFunc(time.Second*10, func() {
		app.Stop() // Stop 关闭服务
	})

	if err := app.Run(); err != nil {
		log.Printf("%v", err)
	}

	logger.Println("app service has exited")
}
