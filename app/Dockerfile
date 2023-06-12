FROM golang:1.18 as build

# 作業ディレクトリを設定
WORKDIR /app

# /appにgo.modをコピーしgo modをダウンロード
COPY . .

# 作業ディレクトリを設定
WORKDIR /app/app

RUN go mod download

# go modのダウンロード、Goアプリのビルド
RUN go build -v main.go

# ポートを公開
EXPOSE 8000

# コンテナ起動時に実行するコマンドを指定
CMD ["go", "run", "main.go"]