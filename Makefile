GOPATH=`pwd`/.go

all: binary container

binary:
	GOPATH=$(GOPATH) go get -d
	GOPATH=$(GOPATH) CGO_ENABLED=0 go build

container:
	docker build -t lew21/waitfor .
