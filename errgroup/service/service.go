package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// App 一个Server生命周期托管对象。
type App struct {
	ctx      context.Context
	cancel   func()
	logger   *log.Logger
	servers  []*http.Server
	errGroup *errgroup.Group
}

// AppOption 让 AppOption包含APP结构
type AppOption func(app *App)

// NewApp 返回一个App对象。
func NewApp(opts ...AppOption) *App {
	app := &App{
		ctx:     context.Background(),       //创建一个根Context
		logger:  log.Default(),              //日志
		servers: make([]*http.Server, 0, 4), //创建容量为4的http server切片
	}

	for _, opt := range opts {
		opt(app)
	}
	//通过 context.WithCancel(ctx), 设定返回的cancel 方法，来级联取消其他 child context 的 goroutine
	app.ctx, app.cancel = context.WithCancel(app.ctx)
	//使用errgroup，进行goroutine的取消
	app.errGroup, app.ctx = errgroup.WithContext(app.ctx)
	return app
}

// Run 开启服务
func (a *App) Run() error {
	a.startListenSystemSignal() // 开启系统信号监听
	a.startServers()            // 开启托管的服务
	return a.errGroup.Wait()    // 等待errGroup的结束信号
}

// Stop 关闭服务
func (a *App) Stop() error {
	a.logger.Println("begin to stop the app...")
	a.cancel()
	return a.errGroup.Wait()
}

// 开启托管的服务
func (a *App) startServers() {
	// 为了处理闭包问题，包一层。
	wrapServerStart := func(srv *http.Server) func() error {
		go func() {
			<-a.ctx.Done() //阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close
			a.logger.Printf("begin to shutdown the server: %v\n", srv.Addr)
			srv.Shutdown(a.ctx) //关闭http server
		}()

		return func() error {
			a.logger.Printf("start the server: %v\n", srv.Addr)
			return errors.WithMessagef(srv.ListenAndServe(), "server is exit: %s", srv.Addr)
		}
	}

	for _, srv := range a.servers {
		a.errGroup.Go(wrapServerStart(srv))
	}
}

// 开启系统信号监听
func (a *App) startListenSystemSignal() {
	a.errGroup.Go(func() error {
		signalCh := make(chan os.Signal, 1) //创建buffer 为1的chan
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-a.ctx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
			signal.Stop(signalCh)
			close(signalCh)
			return a.ctx.Err()
		case signal := <-signalCh: //kill -9 或 其他终止
			return errors.Errorf("receive linux's signal: %v", signal)
		}
	})
}

// WithServer 添加 s1 s2 s3 http server到 APP 对象
func WithServer(server *http.Server) AppOption {
	return func(app *App) {
		app.servers = append(app.servers, server)
	}
}

// WithLog 添加 app 日志
func WithLog(logger *log.Logger) AppOption {
	return func(app *App) {
		app.logger = logger
	}
}
