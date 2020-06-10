/**
* @Author: sui_liut@163.com
* @Date: 2020/5/9 16:05
 */

package controller

import (
	"github.com/docker/docker/client"
)

var DefaultAddr = "unix:///var/run/docker.sock"
var DefaultVersion = client.DefaultVersion

func GetDockerClient()  *client.Client{
	cli, err := client.NewClient(DefaultAddr, DefaultVersion, nil, nil)
	if err != nil {
		panic(err)
	}
	return cli
}