/**
* @Author: sui_liut@163.com
* @Date: 2020/5/14 17:05
 */

package routers

import (
	"containerService/controller"
	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func Router(app *iris.Application) {
	app.Get("/hello", func(ctx context.Context) {
		ctx.JSON(iris.Map{"msg":"ok"})
		return
	})
		container := app.Party("/containers")
	{
		container.Get("/json", (new(controller.ContainerCtl)).List)
		container.Post("/{id}/start", (new(controller.ContainerCtl)).Start)
		container.Post("/{id}/stop", (new(controller.ContainerCtl)).Stop)
		container.Post("/{id}/restart", (new(controller.ContainerCtl)).Restart)
		container.Post("/{id}/kill", (new(controller.ContainerCtl)).Kill)
		container.Post("/{id}/pause", (new(controller.ContainerCtl)).Pause)
		container.Post("/{id}/unpause", (new(controller.ContainerCtl)).Unpause)
		container.Delete("/{id}", (new(controller.ContainerCtl)).Remove)


		container.Get("/{id}/logs", (new(controller.ContainerCtl)).Logs)
		container.Post("/prune", (new(controller.ContainerCtl)).Prune)
		container.Get("/{id}/json", (new(controller.ContainerCtl)).Inspect)
		container.Get("/{id}/top", (new(controller.ContainerCtl)).PS)
		container.Get("/{id}/export", (new(controller.ContainerCtl)).Export)
		container.Get("/{id}/stats", (new(controller.ContainerCtl)).Stats)
	}
}