# Goの公式イメージをベースにする
FROM golang:1.20

# 作業ディレクトリを作成
WORKDIR /app

# 必要なファイルをコピー
COPY main.go .

# Goのビルド（バイナリを作成）
RUN go build -o app main.go

# コンテナのポートを開放
EXPOSE 8080

# 実行するコマンド
CMD ["./app"]
