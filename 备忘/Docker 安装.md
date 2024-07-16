# Docker 安装
## 在 CentOS7 中安装 Docker
1. 配置 Docker yum 源

打开[阿里云 Docker CE 镜像网站](https://developer.aliyun.com/mirror/docker-ce)

获取到 yum 源地址：```https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo```
并安装到 yum 目录
```bash
wget -O /etc/yum.repos.d/docker-ce.repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```
刷新并验证
```bash
sudo yum makecache fast
suod yum repolist
```

2. 安装 docker
```bash
# 安装
sudo yum install docker-ce -y
# 查看
sudo docker version
```

999. 重要事项

关闭防火墙
```bash
sudo systemctl stop firewalld
sudo systemctl disable firewalld
firewall-cmd --state
# not running
```

关闭 SELinux
```bash
vim /etc/selinux/config

SELINUXTYPE=disabled

sestatus
# SELinux status: disabled
```
> 如果关闭 selinux 之后 centos 启动失败
> 重启系统后在内核选择时按 e 进入编辑模式，找到 linux 或 linux16 开头行，并在行尾添加 selinux=0

如果安装完 docker 之后 ```Chain FORWARD (policy DROP)```
```bash
iptables -nL
# Chain FORWARD (policy DROP)
```
修改文件 ```vim /usr/lib/systemd/system/docker.service```
```bash
# 删除后面内容
ExecStart=/usr/bin/dockerd # -H fd:// --containerd=/run/containerd/containerd.sock

# 添加下面内容
ExecStartPost=/sbin/iptables -P FORWARD ACCEPT
```
重启 Docker
```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```