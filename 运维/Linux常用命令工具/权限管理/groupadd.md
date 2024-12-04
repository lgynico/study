# groupadd

添加组
格式：`groupadd [options] <group>`

## 常见用法
```bash
# 添加组
groupadd <group>

# 指定组 ID
group -g <GID> <group>
```

## 扩展
### 组信息文件 `/etc/group`
每行为一个组：`组名:组密码点位符:GID:附加用户`
### 组密码文件 `/etc/gshadow`
每行为一个组：`组名:组密码:管理员:附加用户`