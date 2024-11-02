# ln

创建链接文件 ( link )
格式：`ln [options] <src> <dest>`


## 常见用法
```bash
# 创建硬链接
ln <link_flie> /path/to/<dest_file>

# 创建软链接
ln -s <link_flie> /path/to/<dest_file>
```


## 注意事项
### 软硬链接的区别
||硬链接|软链接|
|:-|:-|:-|
|可链接目录||:white_check_mark:|
|可以跨分区||:white_check_mark:|
|源文件存在关联||:white_check_mark:|