FROM golang:1.18 as build

# 作業ディレクトリを設定
WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /hackathon-server

CMD ["/hackathon-server"]