# alias

命令别名操作
格式：`alias [options]`


## 常见用法
```bash
# 列出所有别名
alias
alias -p

# 创建别名
alias <name>='value'

```

## 扩展
### 永久别名
将别名写入`~/.bashrc`文件中，重启终端后生效
```bash
cat ~/.bashrc

# ...
alias <name>='value'
# ...

source ~/.bashrc
```