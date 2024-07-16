# CentOS7 设置 yum 源

1. 备份 yum 文件
```bash
sudo mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
```

2. 下载阿里云 yum 源
```bash
wget -O /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo

# or
curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo
```

3. (可选) 手动配置 yum 源
```bash
sudo vi /etc/yum.repos.d/CentOS-Base.repo
```
注释所有的 mirrorlist 行，并添加下面的行
```bash
baseurl=http://mirrors.aliyun.com/centos/7/os/x86_64/
```

4. 让 yum 源生效
```bash
sudo yum clean all

sudo yum makecache
```

5. 查看 yum 源是否生效
```bash
sudo yum repolist
```