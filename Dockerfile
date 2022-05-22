FROM golang:alpine

RUN mkdir /build
WORKDIR /build

RUN go env -w GO111MODULE=off
RUN go get github.com/keithlemon/golang-docker-demo/
RUN cd /build && git clone https://github.com/keithlemon/golang-docker-demo

RUN cd /build/golang-docker-demo/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/golang-docker-demo/main/main" ]