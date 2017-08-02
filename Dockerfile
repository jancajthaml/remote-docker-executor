FROM golang:1.7

MAINTAINER Jan Cajthaml <jan.cajthaml@gmail.com>

ENV \
	GOPATH=/go \
	PATH=$GOPATH/bin:$PATH \
	GOOS=linux \
	GOARCH=amd64

RUN go get github.com/tools/godep \
		   github.com/labstack/echo && \
	\
	mkdir -p /go/src/github.com/jancajthaml/remote-docker-executor && \
	curl https://get.docker.com/builds/Linux/x86_64/docker-1.12.1.tgz | tar zx -C /opt && \
	ln -s /opt/docker/docker /usr/bin/docker && docker -v

COPY main.go \
	 bash.go \
	 docker.go \
	 containers.go \
	 \
	 /go/src/github.com/jancajthaml/remote-docker-executor/

WORKDIR /go/src/github.com/jancajthaml/remote-docker-executor

RUN go build \
	\
	main.go \
	bash.go \
	docker.go \
	containers.go 

EXPOSE 8181

CMD ./main