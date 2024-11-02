# cp

拷贝文件 ( copy )
格式：`cp [options] src dest`


## 常见用法
```bash
# 拷贝文件
cp <src_file> <dest_file>
cp <src_file> /path/to/<dest_file>
cp <scr_file> /path/to/

# 拷贝目录
cp -r <src_dir> <dest_dir>
cp -r <src_dir> /path/to/<dest_dir>

# 保留属性
cp -p <src_file> <dest_file>
cp -pr <src_dir> <dest_dir>

# 显示进度
cp -v <src_file> <dest_file>

```