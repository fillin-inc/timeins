# マルチステージビルド対応
# 開発環境とプロダクション環境の両方に対応

# 開発環境
FROM golang:1.20-alpine AS dev
WORKDIR /app

# 開発に必要なツールをインストール
RUN apk add --no-cache git make

# golangci-lintをインストール
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Go modulesの依存関係をダウンロード
COPY go.mod ./
RUN go mod download

# ソースコードをコピー
COPY . .

# デフォルトコマンド
CMD ["make", "test"]

# プロダクション環境用のビルドステージ
FROM golang:1.20-alpine AS builder
WORKDIR /app

# ビルドに必要な最小限のツールをインストール
RUN apk add --no-cache git

# Go modulesの依存関係をダウンロード
COPY go.mod ./
RUN go mod download

# ソースコードをコピー
COPY . .

# バイナリをビルド（静的リンク）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# プロダクション環境用の最終イメージ
FROM alpine:latest AS prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# ビルダーステージからバイナリをコピー
COPY --from=builder /app/main .

CMD ["./main"]
