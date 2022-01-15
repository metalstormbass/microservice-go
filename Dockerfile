
FROM golang:1.15-alpine3.12 as builder
WORKDIR $GOPATH/src/github.com/metalstormbass/microservices-go
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp microservice-go /

FROM alpine:3.12
COPY --from=builder / /
CMD ["/microservice-go"]