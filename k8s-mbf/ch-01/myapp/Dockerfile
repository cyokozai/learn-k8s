# 最新のGoの公式イメージをベースにする
FROM golang:latest

# 環境変数を設定
# OSをLinuxに設定
ENV GOOS=linux
# アーキテクチャをamd64に設定
ENV GOARCH=amd64

# 作業ディレクトリを作成
WORKDIR /app

# 必要なファイルをコピー
COPY testapp.go .

# Goのビルド（バイナリを作成）
RUN go build -o app testapp.go && chmod +x app

# コンテナのポートを開放
EXPOSE 8080

# 実行するコマンド
CMD ["./app"]
