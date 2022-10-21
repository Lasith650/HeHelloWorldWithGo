FROM golang:1.17-alpine as build-env

ENV APP_NAME HelloWorldWithGo
ENV CMD_PATH main.go

COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

FROM alpine:3.14

ENV APP_NAME HelloWorldWithGo

COPY --from=build-env /$APP_NAME .

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd

EXPOSE 8081

USER nobody

CMD ./$APP_NAME