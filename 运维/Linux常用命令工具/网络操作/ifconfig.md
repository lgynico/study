# ifconfig

查看网卡信息
格式：`ifconfig`

## 常见用法
```bash
# 查看网上信息
ifconfig

# 查看指定网卡信息
ifconfig eth0

# 启用指定网卡
ifconfig eth0 up

# 禁用指定网卡
ifconfig eth0 down

# 修改网卡ip
ifconfig eth0 192.168.0.10 netmask 255.255.255.0 broadcast 192.168.0.255

```


## 扩展
### 安装 ifconfig
```bash
yum install net-tools -y

apt install net-tools -y
```

### ifconfig 字段含义
|网卡|含义|备注|
|:-|:-|:-|
|eth0|||
|ens32|||
|lo|||

|字段|含义|备注|
|:-|:-|:-|
|flags|||
|mtu|||

### 网卡关联文件 `/etc/sysconfig/network-scripts/ifcfg-eth0`
虚拟机为 `/etc/sysconfig/network-scripts/ifcfg-ens32`
```bash
cat /etc/sysconfig/network-scripts/ifcfg-ens32

TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="none"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens32"
UUID="984d4080-fe38-4118-bec8-df36bbff36cb"
DEVICE="ens32"
ONBOOT="yes"
IPADDR="192.168.0.10"
PREFIX="24"
GATEWAY="192.168.0.254"
DNS1="223.5.5.5"
IPV6_PRIVACY="no"
```

###