# nmcli

网卡操作
格式：`nmcli [option]`
> 只有 CentOS 系统有？

## 常见用法
```bash
# 查看所有网络设备
nmcli

# 激活网卡
nmcli connection up <网卡名>

# 停用网卡
nmcli connection down <网卡名>

# 重启网卡
nmcli connection reload <网卡名>

# 查看所有配置的网卡
nmcli connection show

# 修改网卡信息（红帽认证）
nmcli connection modify <网卡名> ipv4.method manual ipv4.addresses <ip/mask> connection.autoconnect yes

```

## 扩展