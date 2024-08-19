# Docker 常用软件安装

## MySQL 5.7
1. 拉取镜像
```sh
sudo docker pull mysql:5.7
```
2. 创建映射目录
```sh
sudo mkdir /usr/local/docker_data/mysql/{data,conf,logs}
```
3. 启动容器
```sh
sudo docker run --name mysql5.7 \
    -p 3306:3306 \
    -v /usr/local/docker_data/mysql/data:/var/lib/mysql \
    -v /usr/local/docker_data/mysql/conf:/etc/mysql/conf.d \
    -v /usr/local/docker_data/mysql/logs:/var/log/mysql \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -d mysql:5.7
```