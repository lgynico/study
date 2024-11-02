# free

查看内存信息
格式：`free [options]`


## 常见用法
```bash
# 查看内存信息
free

# 以千字节为单位
free -k

# 以兆字节为单位
free -m

# 以人可读的方式查看内存信息
free -h

# 显示内存使用率
free -w

```


## 扩展
### free 字段含义
行
|字段|含义|备注|
|:-|:-|:-|
|**Mem**|内存||
|**Swap**|交换分区||

列
|total|used|free|shared|buff/cache|available|
|:-|:-|:-|:-|:-|:-|
|总大小|已使用|剩余|共享|缓冲和缓存|可用|


### free 与 top 命令

### 内存关联文件 `/proc/meminfo`
```bash
cat /proc/meminfo

MemTotal:        1863032 kB
MemFree:         1088316 kB
MemAvailable:    1489896 kB
Buffers:            2108 kB
Cached:           522864 kB
SwapCached:            0 kB
Active:           499616 kB
Inactive:          86924 kB
Active(anon):      61944 kB
Inactive(anon):     9220 kB
Active(file):     437672 kB
Inactive(file):    77704 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:       2097148 kB
SwapFree:        2097148 kB
Dirty:                 0 kB
Writeback:             0 kB
AnonPages:         61776 kB
Mapped:            23132 kB
Shmem:              9596 kB
Slab:              72268 kB
SReclaimable:      34580 kB
SUnreclaim:        37688 kB
KernelStack:        3920 kB
PageTables:         3668 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:     3028664 kB
Committed_AS:     284000 kB
VmallocTotal:   34359738367 kB
VmallocUsed:      183408 kB
VmallocChunk:   34359310332 kB
Percpu:            33280 kB
HardwareCorrupted:     0 kB
AnonHugePages:      6144 kB
CmaTotal:              0 kB
CmaFree:               0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
DirectMap4k:       75648 kB
DirectMap2M:     2021376 kB
DirectMap1G:           0 kB
```

### meminfo 字段含义