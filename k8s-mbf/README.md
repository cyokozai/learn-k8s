# つくって、壊して、直して学ぶ Kubernetes 入門

## 目次

- [つくって、壊して、直して学ぶ Kubernetes 入門](#つくって壊して直して学ぶ-kubernetes-入門)
  - [目次](#目次)
  - [Chapter 1.1 作ってみよう Kubernetes | Doker コンテナを作ってみる](#chapter-11-作ってみよう-kubernetes--doker-コンテナを作ってみる)
    - [コンテナ](#コンテナ)
    - [なぜコンテナを使うのか？](#なぜコンテナを使うのか)
      - [仮想マシン (VM) よりも高速にアプリケーションを起動できるようになった](#仮想マシン-vm-よりも高速にアプリケーションを起動できるようになった)
      - [マイクロサービスアーキテクチャとコンテナの相性が良い](#マイクロサービスアーキテクチャとコンテナの相性が良い)
    - [Docker](#docker)
      - [Docker のインストール (Ubuntu)](#docker-のインストール-ubuntu)
      - [Docker の基本コマンド](#docker-の基本コマンド)
    - [Docker image と　Dockerfile](#docker-image-とdockerfile)
  - [Chapter 1.2 作ってみよう Kubernetes | Kubernetes クラスタを作ってみる](#chapter-12-作ってみよう-kubernetes--kubernetes-クラスタを作ってみる)

## Chapter 1.1 作ってみよう Kubernetes | Doker コンテナを作ってみる

### コンテナ

コンテナは**アプリケーションとその実行環境をパッケージ化したもの**である。  
コンテナは名前空間 (namespace) や cgroup により通常のプロセスから隔離された仮想環境である。  
したがって、通常コンテナ内部のプロセスを外部から参照することはできない。  
ただし、`docker exec` コマンドや `nsenter` コマンドを用いると、ホストからコンテナのプロセス空間に対してアクセスすることが可能である。  

### なぜコンテナを使うのか？

Docker の登場以来、コンテナ仮想化技術を採用した開発・運用は盛んに行われている。  
ではなぜコンテナを選ぶケースが増えているのか。  
ここでは主に2つの理由について解説する。  

#### 仮想マシン (VM) よりも高速にアプリケーションを起動できるようになった

コンテナとよく比較される VM は、ハードウェアや OS を含めて仮想化する技術である。  
コンテナはそれ自体は OS を含まず、ホストの Kernel を共有する。  
それゆえ、リソースの消費が少なく起動時間も早いことからコンテナの採用率が増加していると考えられる。  

![image1](./images/vmcontainer.png)

#### マイクロサービスアーキテクチャとコンテナの相性が良い

[マイクロサービスアーキテクチャ (Microservice Architecture)](https://knowledge.sakura.ad.jp/20167/) とは、一般に次のようなものを指す。  

1. 個々のマイクロサービスはそれぞれ独立したプロセスとして動作する
1. 各マイクロサービスは主にネットワーク経由で通信して所定のタスクを処理する
1. 各マイクロサービスはほかのマイクロサービスに依存せず起動でき、独立してデプロイやアップデートが可能

マイクロサービスの採用により、顧客に対して柔軟かつ高速にサービスの価値を提供できるようになる。  
したがって、マイクロサービスの性能を十分に発揮できるようなインフラとして、コンテナが採用される背景がある。  

![image2](./images/monorisandms.png)

### Docker

Docker はコンテナ仮想化技術の一つであり、**コンテナを作成・実行・管理するためのツール**である。  
Dockerfile, Docker image, Docker Hub とこれら一連の操作を提供する Docker CLI により、ユーザにとって一連のライフサイクルを扱いやすくしてくれている。  
Docker を使うことで、どの OS や環境でコンテナを実行しても必ず同じ動作をすることが保証される。

#### Docker のインストール (Ubuntu)

[Docker CE の入手](https://docs.docker.jp/engine/installation/linux/docker-ce/ubuntu.html)を参照。

#### Docker の基本コマンド

<details><summary>基本コマンド一覧を開く</summary><div>

- イメージを取得  
  (公式の nginx コンテナイメージをダウンロード)

    ```shell
    docker pull nginx
    ```

- コンテナを起動  
  (nginx をバックグラウンドで起動し、8080 でアクセス可能にする)

    ```shell
    docker run -d -p 8080:80 nginx
    ```

- 実行中のコンテナを確認

    ```shell
    docker ps
    ```

- コンテナを停止

    ```shell
    docker stop <コンテナID>
    ```

- コンテナを削除

    ```shell
    docker rm <コンテナID>
    ```

- イメージを削除

    ```shell
    docker rmi <イメージID>
    ```

</div></details>

### Docker image と　Dockerfile

- Docker image
  Docker でコンテナを作成するには、コンテナの元となるイメージ (image) を取得する必要があります。  
  例えば、`nginx` のコンテナをデプロイする場面を考える。  
  `nginx` のイメージは Nginx 公式が Docker Hub にイメージを公開しているものを使用する。  
  ```shell
  $ docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d nginx
  ```

## Chapter 1.2 作ってみよう Kubernetes | Kubernetes クラスタを作ってみる
