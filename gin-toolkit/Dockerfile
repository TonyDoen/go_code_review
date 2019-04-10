FROM golang:alpine as builder

WORKDIR $GOPATH/src/gin-toolkit/
COPY . $GOPATH/src/gin-toolkit/
RUN go build -o /usr/local/bin/gin-toolkit main.go

FROM alpine:latest
ENV GIN_SERVICE="gin-toolkit"
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
RUN apk add --no-cache ca-certificates tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /usr/local/bin/godemo /usr/local/bin/

CMD ["/usr/local/bin/gin-toolkit"]
