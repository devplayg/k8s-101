FROM golang:1.15.2
MAINTAINER devplayg <devplayg@gmail.com>

RUN mkdir /app
WORKDIR /app
RUN go get -u "github.com/devplayg/k8s-101/webserver"
RUN  go build -o main $GOPATH/src/github.com/devplayg/k8s-101/webserver
CMD ["./main"]

EXPOSE 80
