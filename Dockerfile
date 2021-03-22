FROM golang:1.15.2-buster

COPY . /tgnotify 

WORKDIR /tgnotify/cmd 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o tgnotify 

FROM alpine:3.12

COPY --from=0 /tgnotify/cmd/tgnotify /bin/
COPY ./dockerinit /

RUN apk --no-cache add ca-certificates

VOLUME /config
EXPOSE 8333 

CMD ["/bin/tgnotify", "--conf", "/config/config.json"]