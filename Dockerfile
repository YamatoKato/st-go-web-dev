# Build stage
# golang:1.23.0-alpine
# イメージサイズを小さくするためにalpineを使用（git,gcc,bashなどが含まれていない）
FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# -o オプションでapp/cli/main.goをビルドしてmainという名前のバイナリを作成
RUN go build -o main /app/app/cli/main.go

# Run stage
# Goで作成したバイナリはAlpine Linux上で実行可能
FROM alpine:3.17
WORKDIR /app
# ビルドステージで作成したバイナリをコピー
COPY --from=builder /app/main .
# ポート番号を指定
EXPOSE 8080
# バイナリを実行
CMD ["./main"]
