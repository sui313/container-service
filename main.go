/**
* @Author: sui_liut@163.com
* @Date: 2020/5/8 16:35
 */

package main

import (
	"containerService/controller"
	"containerService/routers"
	stdContext "context"
	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	dockerversion := os.Getenv("DockerVersion")
	if dockerversion != "" {
		controller.DefaultVersion = dockerversion
	}
	if port == "" {
		if len(os.Args) == 2 {
			port = os.Args[1]
		}
	}
	if port == "" {
		port = "8888"
	}
	app := iris.New()
	app.Logger().SetLevel("debug")

	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)

	app.Use(recover.New())
	app.Use(logger.New())
	routers.Router(app)
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			// kill -SIGINT XXXX 或 Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill等同于syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()
	app.Run(iris.Addr(":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}
