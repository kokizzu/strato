FROM golang:1.12.7 AS go-builder

ADD . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o strato-example-app ./examples/app

FROM alpine:3.9.4

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=go-builder /build/strato-example-app .

ENTRYPOINT [ "./strato-example-app" ]