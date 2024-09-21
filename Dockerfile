# ベースイメージとしてGoを使用
FROM golang:1.18-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールのキャッシュを利用するためにgo.modとgo.sumをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# 環境変数を設定
ENV GO_ENV=dev

# アプリケーションを実行
CMD ["./main","go", "run", "migrate.go"]