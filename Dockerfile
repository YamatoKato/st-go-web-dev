# デプロイ用コンテナに含めるバイナリを作成するコンテナ
FROM golang:1.22.7-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# バイナリを作成
RUN go build -trimpath -ldflags "-s -w" -o app

#--------------------------------------------------

# デプロイ用のコンテナ(リリース想定用)
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

#--------------------------------------------------

# ローカル開発環境で利用するホットリロード環境
FROM golang:1.22.7-bullseye as dev
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
