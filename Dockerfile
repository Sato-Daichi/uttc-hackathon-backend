FROM golang:1.18 as build

COPY ./main.go /main.go

CMD go run /main.go