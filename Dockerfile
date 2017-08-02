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
	ln -s /opt/docker/docker /usr/bin/docker && docker -v && \
	\
	echo "docker inspect -f '{{range \$p, \$conf := .NetworkSettings.Ports}}{{(index \$conf 0).HostPort}}{{end}}' \$(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0" > /usr/local/bin/get_exposed_port && \
	chmod +x /usr/local/bin/get_exposed_port && \
	\
	echo "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' \$(cat /etc/hosts | tail -1 | cut -f2) 2> /dev/null || echo 0.0.0.0" > /usr/local/bin/get_ip && \
	chmod +x /usr/local/bin/get_ip

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