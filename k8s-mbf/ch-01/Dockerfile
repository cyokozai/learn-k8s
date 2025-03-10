# 最新のGoの公式イメージをベースにする
FROM golang:latest

# 作業ディレクトリを作成
WORKDIR /app

# 必要なファイルをコピー
COPY testapp.go .

# Goのビルド（バイナリを作成）
RUN go build -o app testapp.go

# コンテナのポートを開放
EXPOSE 8080

# 実行するコマンド
CMD ["./app"]
