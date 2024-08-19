# Minikube 安装过程问题
## Docker 无法使用 root 账号
使用 `--force` 参数

## 镜像拉取失败
切换到国内镜像：
1. 执行 `minibube delete`
2. 使用 `--image-mirror-country='cn'` 参数

## 拉取到镜像后一直卡住
虚拟机配置调到 4c8g

## 磁盘空间不足
执行 `minikube delete --all --purge`