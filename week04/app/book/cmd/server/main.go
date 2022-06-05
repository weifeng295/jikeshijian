package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"week04/app/book/internal/conf"
	"week04/app/book/internal/pkg"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	defer cancel()

	// listen os signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// waiting for signals
	go func() {
		<-sigs
		fmt.Println("[BOOK API] linux cancel")
		cancel()
	}()

	// override default config
	config := conf.New(
		conf.WithMode(gin.ReleaseMode),
		conf.WithHTTPAddr(":8080"),
		conf.WithGRPCAddr(":2223"),
		conf.WithDatabase(conf.DatabaseOptions{
			Driver:     "mysql",
			DataSource: "root:Sunday.2020!@@tcp(127.0.0.1:3306)/hello?parseTime=True",
		}))

	// run app
	pkg.Logo()
	app := initApp(group, config)
	app.HttpServer.Serve(ctx)
	app.GRPCServer.Serve(ctx)

	// waiting for error
	if err := group.Wait(); err != nil {
		fmt.Println("[BOOK API] error:", err)
	}
}
