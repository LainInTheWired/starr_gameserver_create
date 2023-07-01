# image取得(容量が小さいalpineを選択)
FROM golang:1.20.5-alpine3.17

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/goSouce


