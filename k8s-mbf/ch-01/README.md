# Chapter 1.1 作ってみよう Kubernetes

- [Chapter 1.1 作ってみよう Kubernetes](#chapter-11-作ってみよう-kubernetes)
  - [Dokerfile](#dokerfile)

## Dokerfile

- `docker build`

  ```shell
  docker build ./ch-01/myapp/ --tag testapp:1.0.0
  ```

- `docker images` を実行して自作イメージが作成されたことを確認する  

  ```shell
  $ docker images
  REPOSITORY       TAG       IMAGE ID       CREATED        SIZE
  testapp          1.0.0     ec0cb3683811   20 hours ago   938MB
  ```
