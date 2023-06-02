FROM golang:1.18 as build

COPY . .

CMD go run ./app/main.go