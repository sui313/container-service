/**
* @Author: sui_liut@163.com
* @Date: 2020/5/9 15:53
 */

package controller

import (
	stdContext "context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	context "github.com/kataras/iris/v12/context"
	"io"
	"log"
	"net/http"
	"time"
)

type ContainerCtl struct{}

func (this *ContainerCtl) ReturnJSon(ctx context.Context, code int, msg string, data ...interface{}) {
	ctx.StatusCode(code)
	ctx.JSON(CommonResp{msg, data})
	return
}

// List containers
func (this *ContainerCtl) List(ctx context.Context) {

	var req ContainerListReq

	if err := ctx.ReadQuery(&req); err != nil && !context.IsErrPath(err) {
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rrr := ctx.FormValues()
	req.Filters, _ = getmap(rrr, "filters")

	cli := GetDockerClient()
	var opt types.ContainerListOptions
	opt.All = req.All
	opt.Limit = req.Limit
	opt.Size = req.Size
	opt.Filters = filters.NewArgs()
	for k, v := range req.Filters {
		opt.Filters.Add(k, v)
	}

	container, err := cli.ContainerList(stdContext.Background(), opt)
	if err != nil {
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok", container)
	return
}

// Remove a container
func (this *ContainerCtl) Remove(ctx context.Context) {
	var req ContainerRemoveReq
	var id string
	id = ctx.Params().Get("id")

	if err := ctx.ReadForm(&req); err != nil {
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}

	cli := GetDockerClient()
	err := cli.ContainerRemove(stdContext.Background(), id, types.ContainerRemoveOptions{req.V, req.Link, req.Force})
	if err != nil {
		log.Println("ContainerRemove err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Start a container
func (this *ContainerCtl) Start(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	err := cli.ContainerStart(stdContext.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		log.Println("ContainerStart err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Stop a container
func (this *ContainerCtl) Stop(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")
	var t int
	t, _ = ctx.Params().GetInt("t")
	var d *time.Duration
	if t > 0 {
		d = new(time.Duration)
		*d = time.Duration(t) * time.Second
	}
	cli := GetDockerClient()
	err := cli.ContainerStop(stdContext.Background(), id, d)
	if err != nil {
		log.Println("ContainerStop err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Restart a container
func (this *ContainerCtl) Restart(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")
	var t int
	t, _ = ctx.Params().GetInt("t")
	var d *time.Duration
	if t > 0 {
		d = new(time.Duration)
		*d = time.Duration(t) * time.Second
	}
	cli := GetDockerClient()
	err := cli.ContainerRestart(stdContext.Background(), id, d)
	if err != nil {
		log.Println("ContainerRestart err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Kill a container
func (this *ContainerCtl) Kill(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	signal:= ctx.FormValue("signal")

	cli := GetDockerClient()
	err := cli.ContainerKill(stdContext.Background(), id, signal)
	if err != nil {
		log.Println("ContainerKill err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}


// Pause a container
func (this *ContainerCtl) Pause(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	err := cli.ContainerPause(stdContext.Background(), id)
	if err != nil {
		log.Println("ContainerPause err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Unpause a container
func (this *ContainerCtl) Unpause(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	err := cli.ContainerUnpause(stdContext.Background(), id)
	if err != nil {
		log.Println("ContainerUnpause err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok")
	return
}

// Get container logs
func (this *ContainerCtl) Logs(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")
	var req ContainerLogsReq
	if err := ctx.ReadForm(&req); err != nil {
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !req.Stderr && !req.Stdout {
		req.Stderr = true
	}

	cli := GetDockerClient()
	resp, err := cli.ContainerLogs(stdContext.Background(), id, types.ContainerLogsOptions{ShowStdout: req.Stdout,
		ShowStderr: req.Stderr,
		Since:      fmt.Sprintf("%d", req.Since),
		Timestamps: req.Timestamps,
		Follow:     req.Follow,
		Tail:       req.Tail,
		Details:    false})
	if err != nil {
		log.Println("ContainerLogs err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	io.Copy(ctx.ResponseWriter(), resp)
	resp.Close()
	return
}

// Delete stopped containers
func (this *ContainerCtl) Prune(ctx context.Context) {

	req := ctx.URLParams()
	var pruneFilters filters.Args
	for k, v := range req {
		pruneFilters.Add(k, v)
	}

	cli := GetDockerClient()
	resp, err := cli.ContainersPrune(stdContext.Background(), pruneFilters)
	if err != nil {
		log.Println("ContainersPrune err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok", resp)
	return
}

// Inspect a container
func (this *ContainerCtl) Inspect(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	resp, err := cli.ContainerInspect(stdContext.Background(), id)
	if err != nil {
		log.Println("ContainerInspect err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok", resp)
	return
}

// List processes running inside a container
func (this *ContainerCtl) PS(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")
	ps_args := ctx.URLParam("ps_args")

	cli := GetDockerClient()
	resp, err := cli.ContainerTop(stdContext.Background(), id, []string{ps_args})
	if err != nil {
		log.Println("ContainerTop err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok", resp)
	return
}

// Export a container
func (this *ContainerCtl) Export(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	resp, err := cli.ContainerExport(stdContext.Background(), id)
	if err != nil {
		log.Println("ContainerExport err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	io.Copy(ctx.ResponseWriter(), resp)
	resp.Close()
	return
}

// Get container stats based on resource usage
func (this *ContainerCtl) Stats(ctx context.Context) {
	var id string
	id = ctx.Params().Get("id")

	cli := GetDockerClient()
	resp, err := cli.ContainerStats(stdContext.Background(), id, true)
	if err != nil {
		log.Println("ContainerExport err:", err.Error())
		this.ReturnJSon(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.ReturnJSon(ctx, http.StatusOK, "ok", resp)
	return
}
