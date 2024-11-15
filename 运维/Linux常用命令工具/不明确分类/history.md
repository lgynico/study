# history

历史命令
格式：`history [option]`


## 常见用法
```bash
# 查看所有历史命令
history

# 从第 N 条开始显示
history N

# 清空历史命令
history -c

# 删除第 N 条
history -d N

# 保存当前命令到历史记录
history -a

```

## 扩展
### 关联文件 `~/.bash_history`
新命令保存在内存中，退出终端时才会记录

### 历史命令的保存数量
在 `/etc/profile` 中修改 `HISTSIZE` 的值，默认 1000

### history 的快捷操作
|命令|作用|备注|
|:-|:-|:-|
|`!N`|执行第 N 条命令||
|`!cmd`|执行最后一条以 cmd 开头的命令||
|`!!`|执行上一条命令||