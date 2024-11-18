# useradd

添加用户
格式：`useradd [options] <username>`

## 常见用法
```bash
# 创建用户
useradd <username>

# 指定 UID
useradd -u <UID> <username>

# 指定 home 目录
useradd -d <dir> <username>

# 指定用户基本组
useradd -g <GID> <username>

# 指定用户附加组
useradd -G <GID> <username>

# 指定用户的解释器
useradd -s <shell> <username>

# 指定用户描述
useradd -c <desc> <username>
```

## 扩展
### 用户关联文件 `/etc/passwd`
一行表示一个用户，使用 `:` 分隔字段

每个字段的含义
|root|x|0|0|root|/root|/bin/bash|
|-|-|-|-|-|-|-|
|用户名|密码占位符|UID|基本组|用户描述|家目录|解释器|

|UID 范围|作用|描述|
|-|-|-|
|0|超级用户||
|1 - 999|系统伪用户|不能登录系统|
|1000 - 65535|普通用户|由管理员创建|