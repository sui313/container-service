
build:
	CGO_ENABLED="0" go build -ldflags "-s" -o containerService main.go


dockerbuild:
	docker build -t container-service:v1.0 .

dockertag:
	docker tag container-service:v1.0 192.168.3.103:5000/myimages/containerservice:v1.0

dockerpush:
	docker push 192.168.3.103:5000/myimages/containerservice:v1.0	
all: build dockerbuild dockertag dockerpush
