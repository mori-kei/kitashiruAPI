# ベースイメージとしてGoを使用
FROM golang:1.18-alpine

# airのインストール
RUN apk add --no-cache curl && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

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
CMD ["sh", "-c", "go run migrate/migrate.go && air"]