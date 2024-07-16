# Docker 设置国内仓库镜像地址
1. 打开 ```daemon.json``` 文件
```bash
sudo vim /etc/docker/daemon.json
```

2. 然后添加配置
```json
{
    "registry-mirrors": [
        "https://registry.cn-hangzhou.aliyuncs.com",
        "https://mirror.ccs.tencentyun.com",
        "https://05f073ad3c0010ea0f4bc00b7105ec20.mirror.swr.myhuaweicloud.com",
        "https://registry.docker-cn.com",
        "http://hub-mirror.c.163.com",
        "http://f1361db2.m.daocloud.io"
    ]
}
```

3. 重启 Docker
```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```