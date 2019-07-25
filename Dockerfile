FROM golang:1.12-alpine3.9

ENV http_proxy http://192.168.1.219:1087
ENV https_proxy http://192.168.1.219:1087
RUN apk add git
RUN go get github.com/gin-gonic/gin
RUN go get golang.org/x/net/context
RUN go get google.golang.org/grpc
ENV http_proxy=
ENV https_proxy=
