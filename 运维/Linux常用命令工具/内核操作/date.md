# date

日期操作
格式：`date [options] [+format]`

## 常见用法
```bash
# 查看当前时间
date

# 格式化当前时间
date +"format"

# 设置时间
date -s 'YYYY-mm-dd'
date -s 'HH:MM:SS'
date -s 'YYYY-mm-dd HH:MM:SS'

# 显示UTC时间
date -u

# 显示当前时间的年月日
date +%Y%m%d

# 显示当前时间的时分秒
date +%H%M%S
```


## 扩展
### date 格式化常用格式化字符串
|字符串|作用|备注|
|:-|:-|:-|
|`%Y`|年|`%y`|
|`%B`|月|`%b` `%m`|
|`%d`|日||
|`%H`|时||
|`%M`|分||
|`%S`|秒||
|`%s`|Unix Second||
|`%F`|日期|`%x`|
|`%X`|时间||


### 硬件时间
```bash
# 查看硬件时间
clock

# 同步系统与硬件时间
hwclock -s # 系统时间 <-- 硬件时间
hwclock -w # 硬件时间 <-- 系统时间
```

### TODO 同步 HTC 时间