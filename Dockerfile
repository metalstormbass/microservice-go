
FROM golang as builder
WORKDIR $GOPATH/src/github.com/metalstormbass/microservice-go
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp microservice-go /


FROM ubuntu
WORKDIR /microservice-go
COPY --from=builder ./microservice-go .

CMD ./microservice-go