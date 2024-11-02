# tail

显示文件末尾
格式：`tail [options] <file>`


## 常见用法
```bash
# 查看末尾 10 行
tail <file>
tail /path/to/<file>

# 查看末尾 100 行
tail -n 100 <file>
tail -100 <file>

# 实时查看模式
tail -f <file>

```