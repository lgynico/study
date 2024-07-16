# CentOS7 设置静态 IP

1. 修改 ```/etc/sysconfig/network-scripts/ifcfg-ens33``` 文件
```bash
# 配置文件名称：ifcfg-ens33

# 设备类型，通常是以太网
TYPE=Ethernet

# 设备名称，与文件名匹配
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=static  # 设置为 static 表示使用静态 IP 地址，也可以设置为 dhcp（动态获取）
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=ens33  # 网络接口名称
UUID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx  # 设备的唯一标识符，通常由系统生成
DEVICE=ens33  # 与 NAME 相同的网络接口名称
ONBOOT=yes  # 是否在系统启动时激活此网络接口

# 静态 IP 地址设置
IPADDR=192.168.1.100  # IP 地址
NETMASK=255.255.255.0  # 子网掩码
GATEWAY=192.168.1.1  # 默认网关
DNS1=8.8.8.8  # DNS 服务器地址 1
DNS2=8.8.4.4  # DNS 服务器地址 2（可选）

# 其他可能的设置（根据你的需要添加）
# PREFIX=24  # 如果使用 CIDR 表示法，可以指定前缀长度
# IPV6_ADDRESS=xxxx:xxxx:xxxx:xxxx::xx/64  # 如果需要配置 IPv6 地址
# IPV6_GATEWAY=xxxx:xxxx:xxxx:xxxx::xx  # 如果需要配置 IPv6 网关
# ...
```

2. 重启网络服务 ```sudo systemctl restart network```