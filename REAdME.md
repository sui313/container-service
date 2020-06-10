本程序是用来收集k8s中容器的相关信息，程序也是通过k8s部署运行，通过容器运行时挂载node节点的docker.socket来调用docker api.web服务器使用iris搭建。

具体流程如下

![img.png](img/img.png?raw=true)


可以通过makefile来编译构建镜像，通过containerservice.yaml文件可以在k8s中部署运行。

```

build:
	CGO_ENABLED="0" go build -ldflags "-s" -o containerService main.go


dockerbuild:
	docker build -t container-service:v1.0 .

dockertag:
	docker tag container-service:v1.0 192.168.3.103:5000/myimages/containerservice:v1.0

dockerpush:
	docker push 192.168.3.103:5000/myimages/containerservice:v1.0	

all: build dockerbuild dockertag dockerpush


```


```

make all

kubectl create -f containerservice.yaml

```


访问方式为：http://nodeIp:8888/containers/xxxxx


##实现的API接口如下

具体使用方式参考[https://docs.docker.com/engine/api/v1.40/#tag/Container](https://docs.docker.com/engine/api/v1.40/#tag/Container "具体参数使用方式参考")

```

####1.List containers

 GET /containers/json


####2.Start a container

POST /containers/{id}/start


####3.Stop a container

POST /containers/{id}/stop


####4.Restart a container

POST /containers/{id}/restart


####5.Kill a container
POST /containers/{id}/kill


####6.Pause a container

POST /containers//{id}/pause


####7.Unpause a container

POST /containers/{id}/pause


####8.Remove a container

Delete /containers/{id}


####9.Get container logs

GET /containers/{id}/logs


####10.Delete stopped containers

Post /containers/prune


####11.Inspect a container

Get /containers/{id}/json


####12.List processes running inside a container

Get /containers/{id}/top


####13.Export a container

Get /containers/{id}/expor


####14.Get container stats based on resource usage

Get /containers/{id}/stats


```