FROM golang:1.18-alpine

# 必要なツールをインストール
RUN apk add --no-cache curl bash

# airのインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

# golang-migrateのインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールのキャッシュを利用するためにgo.modとgo.sumをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# 環境変数を設定
ENV GO_ENV=dev

# マイグレーションとホットリロードを実行
CMD ["sh", "-c", "go run infrastructure/db.go  && air"]