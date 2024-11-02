# Linux 性能优化实战

## 01 课前准备
 - Ubuntu 18.04

## 02 基础篇：到底应该怎么理解“平均负载”？

### 1. uptime 命令的输出含义

```bash
$ uptime
07:29:45 up 46 min,  4 users,  load average: 0.00, 0.17, 1.24
```
分别为：
 - **07:29:45** 当前时间
 - **up 46 min** 系统运行时间
 - **4 users** 正在登录用户数
 - **load average: 0.00, 0.17, 1.24** 1分钟，5分钟，15分钟的平均负载

###  2. 平均负载的含义
指单位时间内，系统处于 **可运行状态** 和 **不可中断状态** 的平均进程数，也就是 **平均活跃进程数**

如果平均负载为 2 表示：
- 在 2 个 CPU 的系统上刚好完全占用
- 在 4 个 CPU 的系统上有 50% 空闲
- 在 1 个 CPU 的系统上，则有一半进程竞争不到资源

### 3. 平均负载的合理值
一般情况下，**高于 70%** 就需要排查问题
 - 当 1，5，15 的值基本相同，表示负载平稳
 - 当 1 小于 15 说明负载在减小
 - 当 1 大于 15 说明负载在增加，可能是临时的。如果接收或者超过了 CPU 个数表示可能发生了过载问题

### 4. 与 CPU 使用率的关系
> CPU 使用率表示单位时间内 CPU 繁忙情况的统计
- CPU 密集型，大量使用 CPU，两个都会升高
- I/O 密集型，大量 I/O 等待使用平均负载升高，CPU 使用率则不一定
- 大量进程等待 CPU 调度，两者都会升高

### 5. 三个小案例
> 机器：Ubuntu 18.04, 2 CPU, 8G
> 工具：
> - **stress/stress-ng**: Linux 压力测试工具<br>
> - **sysstat**: Linux 性能工具包，包含 mpstat 和 pidstat
> - **mpstat**: 多核 CPU 分析工具
> - **pidstat**: 进程分析工具

1. CPU 密集型进程
```bash
# 使用一个 CPU 模拟 100% 场景
$ stress-ng --cpu 1 --timeout 600
```
```bash
# 查看平均负载
$ watch -d uptime

08:22:24 up  1:39,  4 users,  load average: 1.00, 0.66, 0.32
```
```bash
# 监控所有 CPU，每 5 秒输出一次
$ mpstat -P ALL 5

08:19:21 AM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
08:19:26 AM  all   35.13    0.00    0.13    0.00    0.00    0.13    0.00    0.00    0.00   64.62
08:19:26 AM    0  100.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00    0.00
08:19:26 AM    1    0.00    0.00    0.20    0.00    0.00    0.00    0.00    0.00    0.00   9
```
可以看到有个 CPU 使用达到 100%，但是 iowait 为 0
```bash
# 使用 pidstat 查看进程问题
# 每隔 5 秒输出，重复 1 次
$ mpstat -u 5 1

08:32:28 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
08:32:32 AM     0         3    0.00    0.19    0.00    0.00    0.19     0  kworker/0:0
08:32:32 AM     0         8    0.00    0.19    0.00    0.00    0.19     0  rcu_sched
08:32:32 AM  1000      1700    0.00    0.19    0.00    0.00    0.19     0  sshd
08:32:32 AM     0     10418    0.00    0.19    0.00    0.00    0.19     0  kworker/u4:0
08:32:32 AM     0     10896   95.81    4.38    0.00    0.19  100.00     1  stress-ng-cpu
08:32:32 AM     0     11119    0.19    0.00    0.00    0.00    0.19     0  pidstat

Average:      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
Average:        0         3    0.00    0.19    0.00    0.00    0.19     -  kworker/0:0
Average:        0         8    0.00    0.19    0.00    0.00    0.19     -  rcu_sched
Average:     1000      1700    0.00    0.19    0.00    0.00    0.19     -  sshd
Average:        0     10418    0.00    0.19    0.00    0.00    0.19     -  kworker/u4:0
Average:        0     10896   95.81    4.38    0.00    0.19  100.00     -  stress-ng-cpu
Average:        0     11119    0.19    0.00    0.00    0.00    0.19     -  pidstat
```

2. I/O 密集型进程
```bash
# 模拟 IO 不停地执行 sync
$ stress-ng -i 1 --timeout 600
```
```bash
$ watch -d uptime

08:35:24 up  1:52,  4 users,  load average: 0.97, 0.81, 0.58
```
```bash
$ mpstat -P ALL 5 1

08:37:46 AM  CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
08:37:51 AM  all    0.00    0.00    6.01   24.04    0.00    1.91    0.00    0.00    0.00   68.03
08:37:51 AM    0    0.00    0.00    0.20    0.00    0.00    0.20    0.00    0.00    0.00   99.60
08:37:51 AM    1    0.00    0.00   18.88   75.54    0.00    5.58    0.00    0.00    0.00    0.00

Average:     CPU    %usr   %nice    %sys %iowait    %irq   %soft  %steal  %guest  %gnice   %idle
Average:     all    0.07    0.00    7.85   16.12    0.00    2.26    0.00    0.00    0.00   73.70
Average:       0    0.08    0.00    4.83    0.10    0.00    0.06    0.00    0.00    0.00   94.93
Average:       1    0.08    0.00   13.81   48.01    0.00    6.69    0.00    0.00    0.00   31.41
```
```bash
$ pidstat -u 5 1

08:38:53 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
08:38:58 AM     0       198    0.00    8.37    0.00    0.40    8.37     1  kworker/1:1H
08:38:58 AM     0     11222    0.40   13.15    0.00    2.19   13.55     1  stress-ng-io
08:38:58 AM     0     11697    0.00    0.20    0.00    0.00    0.20     1  pidstat

Average:      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
Average:        0       198    0.00    8.37    0.00    0.40    8.37     -  kworker/1:1H
Average:        0     11222    0.40   13.15    0.00    2.19   13.55     -  stress-ng-io
Average:        0     11697    0.00    0.20    0.00    0.00    0.20     -  pidstat
```

3. 大量进程
```bash
# 开启 8 个进程
$ stress-ng -c 8 --timeout 600
```
```bash
$ uptime

09:20:59 up  2:38,  4 users,  load average: 7.90, 4.79, 2.15
```
```bash
$ pidstat -u 5 1

09:21:15 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:21:20 AM     0         3    0.00    0.20    0.00    0.00    0.20     0  kworker/0:0
09:21:20 AM  1000      1700    0.00    0.20    0.00    0.00    0.20     1  sshd
09:21:20 AM     0     15140   27.29    0.40    0.00   44.02   27.69     0  stress-ng-cpu
09:21:20 AM     0     15141   10.16    0.00    0.00   89.64   10.16     1  stress-ng-cpu
09:21:20 AM     0     15142   33.47    0.20    0.00   26.10   33.67     1  stress-ng-cpu
09:21:20 AM     0     15143   29.48    0.20    0.00   57.77   29.68     0  stress-ng-cpu
09:21:20 AM     0     15144   24.70    0.00    0.00   94.62   24.70     0  stress-ng-cpu
09:21:20 AM     0     15145   17.53    0.00    0.00   82.47   17.53     0  stress-ng-cpu
09:21:20 AM     0     15146   28.29    0.00    0.00   90.24   28.29     1  stress-ng-cpu
09:21:20 AM     0     15147   27.09    0.00    0.00   92.43   27.09     1  stress-ng-cpu
09:21:20 AM     0     15490    0.00    0.40    0.00    1.00    0.40     1  pidstat

Average:      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
Average:        0         3    0.00    0.20    0.00    0.00    0.20     -  kworker/0:0
Average:     1000      1700    0.00    0.20    0.00    0.00    0.20     -  sshd
Average:        0     15140   27.29    0.40    0.00   44.02   27.69     -  stress-ng-cpu
Average:        0     15141   10.16    0.00    0.00   89.64   10.16     -  stress-ng-cpu
Average:        0     15142   33.47    0.20    0.00   26.10   33.67     -  stress-ng-cpu
Average:        0     15143   29.48    0.20    0.00   57.77   29.68     -  stress-ng-cpu
Average:        0     15144   24.70    0.00    0.00   94.62   24.70     -  stress-ng-cpu
Average:        0     15145   17.53    0.00    0.00   82.47   17.53     -  stress-ng-cpu
Average:        0     15146   28.29    0.00    0.00   90.24   28.29     -  stress-ng-cpu
Average:        0     15147   27.09    0.00    0.00   92.43   27.09     -  stress-ng-cpu
Average:        0     15490    0.00    0.40    0.00    1.00    0.40     -  pidstat
```
可以看到有的进程 wait 高于 75%

## 03/04 基础篇：经常说的 CPU 上下文切换是什么意思？
### 1. CPU 上下文的概念
任务执行时必须依赖两个环境：**CPU 寄存器** 和 **程序计数器**

### 2. CPU 上下文切换
切换任务时保存上一个任务的 CPU 上下文，加载下一个任务的 CPU 上下文

### 3. CPU 上下文切换的场景
#### 进程上下文切换
0. 进程上下文包含：虚拟内存、栈、全局变量，内核堆栈、寄存器
1. 进程的内存空间分为内核空间 Ring0 和用户空间 Ring3
2. 进程需要从用户态进入内核态才能访问系统资源，称为系统调用
3. 系统调用会触发两次 CPU 上下文切换：保存 CPU 寄存器用户态的指令，然后加载内核态指令，系统调用结束后，CPU 寄存器恢复到用户态
4. 进程上下文切换比系统调用多一步：先保存进程的虚拟内存、栈然后是内核状态和 CPU 寄存器，加载了下一进程的内核态后，还需要刷新进程的虚拟内存和用户栈
5. 切换时机
> 时间片耗尽
> 所需系统资源不足
> sleep 自主挂起
> 调度高优先级任务
> 硬件中断

#### 线程上下文切换
1. 同一进程内的线程切换
资源共享，只切换线程的私有数据，寄存器

2. 不同进程的线程切换
资源不共享，相当于进程切换

#### 中断上下文切换
1. 中断优先级更高，会打断进程的执行
2. 中断程序没有用户态，不用保存进程资源
3. 中断程序必须短小精悍

### 4. 案例：分析系统上下文切换
> 工具：
> - vmstat: 分析系统内存使用情况，CPU 上下文切换和中断次数
> - sysbench: 多线程基准测试工具

使用 sysbench 模拟多线程调度瓶颈
```bash
# 10 个线程运行 5 分钟
$ sysbench --threads=10 --time=300 threads run
```
通过 vmstat 观察上下文切换
```bash
# 每秒打印 1 次
$ vmstat 1

procs -----------memory---------- ---swap-- -----io---- -system-- ------cpu-----
 r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st
 8  0      0 7430636  40760 552364    0    0     0     0 4277 1462322 10 82  8  0  0
 7  0      0 7430636  40760 552364    0    0     0     0 4205 1462232  7 85  8  0  0
```
关注 4 个字段:
 - **cs** (context switch) 每秒上下文切换次数
 - **in** (interrupt) 每秒中断次数
 - **r** (Running/Runnable) 就绪队列长度，运行中和等待 CPU 的进程数
 - **b** (Blocked) 不可中断睡眠的进程数

可以发现：
- 活跃进程 8 个远远大于 CPU 数量2，说明竞争激烈
- 内存使用率(us+sy)也接近 100%，sy 较高说明显内核占用多
- 中断次数 4000 多，上下文切换高达 140 多万
分析可得系统就绪队列过长，导致了大量的上下文切换，而上下文切换又导致了系统 CPU 的占用率升高

使用 pidstat 查看 CPU 和进程上下文切换情况：
```bash
# -w 输出进程切换指
# -u 输出CPU使用指标
$ pidstat -w -u 1

06:33:37 AM   UID       PID   cswch/s nvcswch/s  Command
06:33:38 AM     0         7      1.00      0.00  ksoftirqd/0
06:33:38 AM     0         8     26.00      0.00  rcu_sched
06:33:38 AM     0        16      2.00      0.00  ksoftirqd/1
06:33:38 AM     0        23      3.00      0.00  kworker/0:1
06:33:38 AM     0        39      5.00      0.00  kworker/1:1
06:33:38 AM  1000      2174     11.00      1.00  sshd
06:33:38 AM     0      3321    247.00      0.00  kworker/u4:1
06:33:38 AM     0      3432      1.00      2.00  vmstat
06:33:38 AM     0      3438      1.00    240.00  pidstat

06:33:38 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
06:33:39 AM     0      3419   14.00  100.00    0.00    0.00  100.00     0  sysbench
06:33:39 AM     0      3438    0.00    1.00    0.00    0.00    1.00     1  pidstat
```
关注两个字段：
- **cswch/s** (voluntary context switches) 每秒自愿上下文切换，无法获取资源，说明系统资源不足
- **nvcswch/s** (non-voluntary context switches) 每秒非自愿上下文切换，时间片耗尽，说明大量进程在竞争 CPU

```bash
$ pidstat -wt 1

06:35:33 AM   UID      TGID       TID   cswch/s nvcswch/s  Command
06:35:34 AM     0      3432         -      1.00      0.00  vmstat
06:35:34 AM     0         -      3432      1.00      0.00  |__vmstat
06:35:34 AM     0         -      3444   6273.00 140156.00  |__sysbench
06:35:34 AM     0         -      3445   7778.00 141517.00  |__sysbench
06:35:34 AM     0         -      3446   5767.00 134919.00  |__sysbench
06:35:34 AM     0         -      3447   5688.00 186147.00  |__sysbench
06:35:34 AM     0         -      3448   7294.00 121317.00  |__sysbench
06:35:34 AM     0         -      3449   7630.00 122438.00  |__sysbench
06:35:34 AM     0         -      3450   6182.00  99196.00  |__sysbench
06:35:34 AM     0         -      3451  11599.00 149911.00  |__sysbench
06:35:34 AM     0         -      3452   4886.00 120546.00  |__sysbench
06:35:34 AM     0         -      3453   7714.00 123726.00  |__sysbench
06:35:34 AM     0      3454         -      1.00    448.00  pidstat
06:35:34 AM     0         -      3454      1.00    448.00  |__pidstat
```
可以看到 sysbench 的子线程是原罪

观察中断次数变高的情况
```bash
$ watch -d cat /proc/interrupts

           CPU0       CPU1
RES:    2970505    3107342   Rescheduling interrupts
```
RES (重调度中断) 变化最快，表示唤醒空闲的 CPU 来调度新的任务
所以还是因为过多任务的调度问题

### 5. 上下文切换次数的合理值
次数比较稳定的情况下，几百到一万以内算正常
超过一万次，或出现数量级增长，可能出现了性能问题

## 05 | 基础篇：某个应用的CPU使用率居然达到100%，我该怎么办？

### 1. CPU 使用率
1. 操作系统通过频率控制时间片，每一次时钟中断都会让变量 `Jiffies` 加 1
2. 内核的节拍为 CONFIG_HZ=250，用户空间节拍为 USER_HZ=100
```bash
$ grep 'CONFIG_HZ=' /boot/config-$(uname -r)
CONFIG_HZ=250
```
3. /proc 为系统内部状态信息，/proc/stat 为系统 CPU 和任务统计信息
```bash
# 列出 CPU 的节拍数(USER_HZ)
$ cat /proc/stat | grep ^cpu

cpu  664 190 1696 2662188 498 0 1260 0 0 0
cpu0 418 75 843 1331765 293 0 925 0 0 0
cpu1 246 115 853 1330423 204 0 334 0 0 0
```
4. 通过 `man proc` 查看每一列的含义
   - **user** (us) 用户态 CPU 时间，不包括 nice
   - **nice** (ni) 低优先级 ( 1 - 19 ) 用户态 CPU 时间
   - **system** (sys) 内核态 CPU 时间
   - **idle** (id) 空闲时间，不包括 iowait
   - **iowait** (wa) 等待 IO 的 CPU 时间
   - **irq** (hi) 硬中断的 CPU 时间
   - **softirq** (si) 软中断的 CPU 时间
   - **steal** (st) 在虚拟机中，被其它虚拟机占用的 CPU 时间
   - **guest** (guest) 运行虚拟机的 CPU 时间
   - **guest_nice** (gnice) 以低优先级运行虚拟机的时间

5. CPU 使用率 = 1 - ( 空闲时间 / 总 CPU 时间 )
6. 平均 CPU 使用率 = 1 - ( ( 空闲时间new - 空闲时间old ) / ( 总CPU时间new - 总CPU时间old ) )

### 2. 查看 CPU 使用率
1. top 工具

%CPU 开头的行就是 CPU 使用率
```bash
# 默认每 3 秒采样一次，按 1 显示每个 CPU
$ top

top - 09:13:50 up  5:06,  2 users,  load average: 0.00, 0.00, 0.00
Tasks:  98 total,   1 running,  51 sleeping,   0 stopped,   0 zombie
%Cpu0  :  0.0 us,  0.3 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
%Cpu1  :  0.0 us,  0.0 sy,  0.0 ni,100.0 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  8167892 total,  7471512 free,   128528 used,   567852 buff/cache
KiB Swap:  4194300 total,  4194300 free,        0 used.  7792768 avail Mem 

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                           
    1 root      20   0  159544   8816   6672 S   0.0  0.1   0:02.89 systemd                                                                           
    2 root      20   0       0      0      0 S   0.0  0.0   0:00.06 kthreadd                                                                          
    4 root       0 -20       0      0      0 I   0.0  0.0   0:00.00 kworker/0:0H 
```
2. pidstat 工具
   - **%usr** 用户态
   - **%system** 内核态
   - **%guest** 虚拟机
   - **wait** 等待
   - **%CPU** 总共
```bash
# 每秒采样一次，共 5 次
$ pidstat 1 5

09:18:18 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
09:18:19 AM  1000      2198    0.00    0.98    0.00    0.00    0.98     0  sshd

...

Average:      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
Average:     1000      2198    0.00    0.19    0.00    0.00    0.19     -  sshd
```

### 3. CPU 使用率过高排查方法
使用 perf 工具
1. perf top 显示占用 CPU 最多的函数或指令
   - **Overhead** 采样比例
   - **Shared** 共享对象，如内核，进程名等
   - **Object** 动态共享对象类型，[.] 用户空间，[k] 内核空间
   - **Symbol** 符号名，函数名，十六进制地址
```bash
$ perf top

Samples: 319  of event 'cpu-clock', Event count (approx.): 69238270
Overhead  Shared Object       Symbol
  50.24%  [kernel]            [k] __softirqentry_text_start
   8.81%  [kernel]            [k] do_idle
   7.39%  [kernel]            [k] _raw_spin_unlock_irqrestore
```

2. perf record / perf report 进行数据采样和离线分析
```bash
$ pref record # Ctrl + C 停止采样

[ perf record: Woken up 3 times to write data ]
[ perf record: Captured and wrote 0.941 MB perf.data (17750 samples) ]

$ perf report
```

### 4. 案例：找到 CPU 使用率高的原因
> 准备
> 机器：两台 Ubuntu 18.04, 2 CPU, 8GB
> 工具：`apt install docker.io sysstat linux-tools-common apache2-utils`
> - docker: 用来安装服务
> - systat: 看上面
> - perf: 通过采样分析性能的工具
> - ab: HTTP 性能测试工具

在第一台机器启动两个服务
```bash
$ docker run --name nginx -p 10000:80 -itd feisky/nginx
$ docker run --name phpfpm -itd --network container:nginx feisky/php-fpm
```
> 这里遇到了 docker 无法下载的问题，使用镜像进行加速
> ```bash
> $ docker pull 0df7nwfc.mirror.aliyuncs.com:feisky/nginx
> $ docker pull 0df7nwfc.mirror.aliyuncs.com:feisky/php-fpm
> ```

试试能不能访问
```bash
$ curl http://192.168.1.15:10000
$ It works!
```

在第二台机器使用 ab 测试 nginx 的性能
```bash
# 10 个并发请求 100 次
$ ab -c 10 -n 100 http://192.168.1.15:10000/

This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 192.168.1.15 (be patient).....done

...

Requests per second:    29.46 [#/sec] (mean)
Time per request:       339.431 [ms] (mean)

```
可以看到每秒请求只有 29.46

继续 10000 个请求，然后回到第一台查看 top
```bash
$ ab -c 10 -n 10000 http://192.168.1.15:10000/
```
```bash
# top 然后按 1
$ top

...
%Cpu0  : 52.0 us,  9.8 sy,  0.0 ni,  0.0 id,  0.0 wa,  0.0 hi, 38.2 si,  0.0 st
%Cpu1  : 29.2 us, 47.2 sy,  0.0 ni,  0.0 id,  0.0 wa,  0.0 hi, 23.6 si,  0.0 st
...

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                             
 2811 daemon    20   0  336696  12872   5196 R  40.4  0.2   0:32.58 php-fpm                                                             
 2808 daemon    20   0  336696  12872   5196 R  37.4  0.2   0:30.20 php-fpm                                                             
 2809 daemon    20   0  336696  12872   5196 R  37.1  0.2   0:28.62 php-fpm                                                             
 2812 daemon    20   0  336696  12932   5256 R  36.1  0.2   0:30.35 php-fpm                                                             
 2810 daemon    20   0  336696  12872   5196 R  30.5  0.2   0:30.11 php-fpm     
```
可以看到 php-fpm 的 CPU 使用率总合接近 200%，每个 CPU 的 us 也很高

使用 perf 工具分析 php-pfm 中的问题函数
```bash
# -g 开启调用关系，-p 指定进程
$ perf top -g -p 2612

Samples: 3K of event 'cpu-clock', Event count (approx.): 166990941
  Children      Self  Shared Object     Symbol                                                                                                                                             ▒+   97.73%     0.10%  libc-2.24.so      [.] epoll_wait                                                                                                                                     ◆+   97.27%    92.92%  [kernel]          [k] finish_task_switch                                                                                                                             ▒-   97.13%     0.54%  php-fpm           [.] php_register_internal_extensions                                                                                                               ▒
     7.38% php_register_internal_extensions                                                                                                                                                ▒        php_register_internal_extensions                                                                                                                                                   ▒+   84.80%     0.07%  [kernel]          [k] __schedule                                                                                                                                     ▒
+   80.73%     0.03%  [kernel]          [k] do_syscall_64                                                                                                                                  ▒+   77.08%     0.05%  [kernel]          [k] sys_epoll_wait                                                                                                                                 ▒+    7.34%     0.00%  [unknown]         [k] 0x6cb6258d4c544155                                                                                                                             ▒+    7.34%     0.00%  libc-2.24.so      [.] __libc_start_main                                                                                                                              ▒+    6.89%     0.00%  [kernel]          [k] entry_SYSCALL_64_after_hwframe   
```
> 文章说的可以看到 php-fpm 的调用关系( sqrt 和 add_function )，但是这里我看不到，解决方法是用 perf record 然后把 perf.data 文件拷贝到 docker 里面再用 perf report 分析
> ```bash
> $ perf record -g -p <pid>
> $ docker cp perf.data phpfpm:/tmp
> $ docker exec -i -t phpfpm bash
> $ cd/tmp/
> $ apt-get update && apt-get install -y linux-perf linux-tools procps
> $ perf_4.9 report
> ```


拷贝出 Nginx 应用的源码，并查找调用的方法
```bash
# 从容器拷贝源码
$ docker cp phpfpm:/app .

# 查找 sqrt 方法调用
$ grep sqrt -r app/
app/index.php:  $x += sqrt($x);

# 查找 add_function 方法调用，因为是内置方法，没找到
$ grep add_function -r app/
```

查看源码，发现问题！
```bash
$ cat app/index.php 
<?php
// test only.
$x = 0.0001;
for ($i = 0; $i <= 1000000; $i++) {
  $x += sqrt($x);
}

echo "It works!"
```

解决问题，使用修复后的镜像
```bash
$ docker rm -f nginx phpfpm
$ docker run --name nginx -p 10000:50 -itd feisky/nginx:cpu-fix
$ docker run --name phpfpm -itd --network container:nginx feisky/php-fpm:cpu-fix
```
> 又没下下来。。。。只能借用文章的内容了

查看修复后的效果
```bash
$ ab -c 10 -n 10000 http://192.168.1.15:10000/

... 
Complete requests:       10000 
Failed requests:         0 
Total transferred:       1720000  bytes 
HTML transferred:        90000  bytes 
Requests  per   second :     2237.04  [# / sec] (mean) 
Time   per  request:        4.470  [ms] (mean) 
Time   per  request:        0.447  [ms] (mean, across  all  concurrent requests) 
Transfer rate:           375.75  [Kbytes / sec] received 
... 
```

## 06 | 案例篇：系统的 CPU 使用率很高，但为啥却找不到高 CPU 的应用？
> 准备：
> 机器：两台 Ubuntu 18.04, 2 CPU, 8GB
> 工具：docker, systat, perf, ab 等
> apt install docker.io sysstat linux-tools-common apache2-utils

终端 1 运行应用
```bash
$ docker run --name nginx -p 10000:80 -itd feisky/nginx:sp
$ docker run --name phpfpm -itd --network container:nginx feisky/php-fpm:sp
```

终端 2 验证
```bash
$ curl http://192.168.1.15:10000/

It works!
```

终端 2 执行 ab 压测
```bash
# 100 个并发请求 1000 次
$ ab -c 100 -n 1000 http://192.168.1.15:10000/

This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
...

Requests per second:    106.09 [#/sec] (mean)
Time per request:       942.555 [ms] (mean)
Time per request:       9.426 [ms] (mean, across all concurrent requests)
Transfer rate:          17.82 [Kbytes/sec] received

```

可以看到每秒只有 106 个请求。

终端 2 继续执行 ab 压测
```bash
# 5 个并发请求 600 秒
$ ab -c 5 -t 600 http://192.168.1.15:10000/
```

终端 1 使用 top 观察 CPU 使用情况
```bash
$ top

top - 04:33:15 up 17 min,  2 users,  load average: 4.79, 3.20, 1.52
Tasks: 127 total,   8 running,  72 sleeping,   0 stopped,   0 zombie
%Cpu(s): 35.2 us, 36.6 sy,  0.0 ni,  6.2 id,  0.0 wa,  0.0 hi, 21.9 si,  0.0 st
KiB Mem :  8167892 total,  7141420 free,   230472 used,   796000 buff/cache
KiB Swap:  4194300 total,  4194300 free,        0 used.  7660952 avail Mem

  PID USER      PR  NI    VIRT    RES    SHR S  %CPU %MEM     TIME+ COMMAND                                                                           
   16 root      20   0       0      0      0 R   8.0  0.0   0:29.68 ksoftirqd/1                                                                       
 2672 systemd+  20   0   33104   3788   2364 R   6.3  0.0   0:24.14 nginx                                                                             
26491 daemon    20   0  336696  15836   8160 S   3.7  0.2   0:02.83 php-fpm
26490 daemon    20   0  336696  15508   7832 S   3.3  0.2   0:02.86 php-fpm
26492 daemon    20   0  336696  15568   7892 R   3.0  0.2   0:02.82 php-fpm                                                                           
26503 daemon    20   0  336696  15836   8160 S   2.7  0.2   0:02.79 php-fpm
26509 daemon    20   0  336696  15836   8160 S   2.7  0.2   0:02.74 php-fpm
 1533 root      20   0 1416416  70948  48064 S   2.3  0.9   0:07.50 dockerd
 2607 root      20   0  712488   9656   7616 S   2.3  0.1   0:08.11 containerd-shim
    5 root      20   0       0      0      0 I   0.3  0.0   0:00.62 kworker/u4:0
    7 root      20   0       0      0      0 S   0.3  0.0   0:00.82 ksoftirqd/0
    8 root      20   0       0      0      0 I   0.3  0.0   0:00.43 rcu_sched
   15 root      rt   0       0      0      0 S   0.3  0.0   0:00.26 migration/1
   38 root      20   0       0      0      0 I   0.3  0.0   0:00.24 kworker/u4:1
 2322 root      20   0       0      0      0 I   0.3  0.0   0:01.88 kworker/0:3
    1 root      20   0  159708   8852   6576 S   0.0  0.1   0:01.29 systemd
    2 root      20   0       0      0      0 S   0.0  0.0   0:00.01 kthreadd
    3 root      20   0       0      0      0 I   0.0  0.0   0:00.00 kworker/0:0
    4 root       0 -20       0      0      0 I   0.0  0.0   0:00.00 kworker/0:0H
    6 root       0 -20       0      0      0 I   0.0  0.0   0:00.00 mm_percpu_wq
    9 root      20   0       0      0      0 I   0.0  0.0   0:00.00 rcu_bh
   10 root      rt   0       0      0      0 S   0.0  0.0   0:00.16 migration/0
   11 root      rt   0       0      0      0 S   0.0  0.0   0:00.00 watchdog/0
   12 root      20   0       0      0      0 S   0.0  0.0   0:00.00 cpuhp/0
   13 root      20   0       0      0      0 S   0.0  0.0   0:00.00 cpuhp/1
```
看到 us 是 35.2 us, sy 是 36.6，si 是 21.9
> 没有看到原文一样的 us 到达 80 的情况，不知道哪里不对
没有看到占用 CPU 特别高的进程，总和也达不到 us 的数值

使用 pidstat 观察一下
```bash
$ pidstat 1

04:31:45 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
04:31:46 AM     0        10    0.00    1.00    0.00    0.00    1.00     0  migration/0
04:31:46 AM     0        16    0.00    8.00    0.00    0.00    8.00     1  ksoftirqd/1
04:31:46 AM     0      1533    0.00    2.00    0.00    0.00    2.00     0  dockerd
04:31:46 AM     0      2607    2.00    0.00    0.00    0.00    2.00     0  containerd-shim
04:31:46 AM   101      2672    1.00    4.00    0.00    4.00    5.00     1  nginx
04:31:46 AM     1     26490    0.00    2.00    0.00    5.00    2.00     0  php-fpm
04:31:46 AM     1     26491    0.00    3.00    0.00    4.00    3.00     1  php-fpm
04:31:46 AM     1     26492    1.00    2.00    0.00    7.00    3.00     0  php-fpm
04:31:46 AM     1     26503    1.00    1.00    0.00    5.00    2.00     0  php-fpm
04:31:46 AM     1     26509    0.00    3.00    0.00    6.00    3.00     1  php-fpm
```
也没发现高 CPU 进程

回头看 top 的情况，发现大多数 php-fpm 进程处于 S 状态，处理 R 状态的是几个 stress 进程
> 原文中可以看到 stress 进程，我看不到... 是不是因为 stress 版本问题

接下来查看其中一个 stress 进程情况
> 我这里没有所以没做这一步
```bash
$ pidstat -p 24344

07:19:16 AM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command


$ ps aux | grep 24344

root     20675  0.0  0.0  13140  1064 pts/0    S+   07:19   0:00 grep --color=auto 24344
```

没有任何输出，说明进程已经不存在，可以用 top 再次看看。
> 由于我实验出来没有看到 stress 进程，所以以下内容无法验证
可以发送 stress 的 pid 一直在变，两个原因：
1. 进程不断崩溃重启
2. 都是短时进程，是应用通过 exec 调用外部命令，top 很难发现


使用 pstree 查看 stress 进程的父子关系，发现是 php-fpm 的子进程
```bash
$ pstree | grep stress

        |                 |         |-3*[php-fpm---sh---stress---stress]
```


接下来把机器 1 的源码拷贝出来，并找出调用 stress 的方法
```bash
$ docker cp phpfpm:/app .

$ grep stress -r app

app/index.php:// fake I/O with stress (via write()/unlink()).
app/index.php:$result = exec("/usr/local/bin/stress -t 1 -d 1 2>&1", $output, $status);


$ cat app/index.php

<?php
// fake I/O with stress (via write()/unlink()).
$result = exec("/usr/local/bin/stress -t 1 -d 1 2>&1", $output, $status);
if (isset($_GET["verbose"]) && $_GET["verbose"]==1 && $status != 0) {
  echo "Server internal error: ";
  print_r($output);
} else {
  echo "It works!";
}

?>
```

请求 verbose=1 之后看看情况
```bash
$ curl http://192.168.1.15:10000?verbose=1

Server internal error: Array
(
    [0] => stress: info: [15223] dispatching hogs: 0 cpu, 0 io, 0 vm, 1 hdd
    [1] => stress: FAIL: [15224] (563) mkstemp failed: Permission denied
    [2] => stress: FAIL: [15223] (394) <-- worker 15224 returned error 1
    [3] => stress: WARN: [15223] (396) now reaping child worker processes
    [4] => stress: FAIL: [15223] (400) kill error: No such process
    [5] => stress: FAIL: [15223] (451) failed run completed in 0s
)
```
可以看到 `Permission denied` 和 `failed run completed`，这是 PHP 调用 stress 的一个 bug，没有创建临时文件的权限

推测：由于权限错误，大量的 stress 进程在启动时初始化失败，进而导致用户 CPU 使用率的升高

使用 perf 工具验证猜想
```bash
$ perf record -g # 等待几秒之后 ctrl + c

$ perf report

Samples: 78K of event 'cpu-clock', Event count (approx.): 19623500000
Overhead  Command          Shared Object            Symbol
   7.07%  ksoftirqd/1      [kernel.kallsyms]        [k] e1000_clean
   5.60%  stress           [kernel.kallsyms]        [k] e1000_clean
   4.96%  stress           libc-2.24.so             [.] random
   4.06%  swapper          [kernel.kallsyms]        [k] native_safe_halt
   3.92%  stress           libc-2.24.so             [.] random_r
   3.64%  sh               libc-2.24.so             [.] strerror_l
   3.61%  stress           libc-2.24.so             [.] strerror_l
   3.38%  stress           [kernel.kallsyms]        [k] exit_to_usermode_loop
   3.29%  nginx            [kernel.kallsyms]        [k] e1000_xmit_frame
   2.37%  stress           stress                   [.] 0x0000000000002f29
   1.93%  stress           [kernel.kallsyms]        [k] __softirqentry_text_start
   1.88%  php-fpm          [kernel.kallsyms]        [k] e1000_clean
   1.81%  stress           [kernel.kallsyms]        [k] _raw_spin_unlock_irqrestore
   1.49%  php-fpm          [kernel.kallsyms]        [k] copy_page
   1.48%  sh               [kernel.kallsyms]        [k] e1000_clean
   1.33%  ksoftirqd/1      [kernel.kallsyms]        [k] e1000_xmit_frame
   1.27%  nginx            [kernel.kallsyms]        [k] e1000_clean
   1.27%  stress           stress                   [.] 0x0000000000002f25
   1.22%  stress           libc-2.24.so             [.] 0x000000000001fd27
   1.18%  sh               libc-2.24.so             [.] 0x000000000001fd27
   1.15%  stress           stress                   [.] 0x0000000000000e80
   1.15%  stress           stress                   [.] 0x0000000000002f14
   1.15%  stress           stress                   [.] 0x0000000000002f09
   1.13%  stress           stress                   [.] 0x0000000000002ef3
   1.13%  stress           stress                   [.] 0x0000000000002f1e
   1.03%  stress           stress                   [.] 0x0000000000002f18
   1.02%  stress           stress                   [.] 0x0000000000002f1b
   0.85%  stress           [kernel.kallsyms]        [k] __do_page_fault
   0.81%  sh               ld-2.24.so               [.] 0x0000000000017287
```
原文中的 stress 占用 CPU 77, 就是它了

结束，清理环境
```bash
$ docker rm -f nginx phpfpm
```


## 07 & 08 | 系统中出现大量不可中断进程和僵尸进程怎么办？
僵尸进程：父进程来不及回收退出的子进程

### 1. 进程状态
top 命令中的进程标识（S）
- R (Running/Runnable) 运行中或者就绪队列中
- D (Disk Sleep) 不可中断睡眠
- Z (Zombie) 僵尸进程
  wait() / waitpid() 等待子进程结束
  child -- SIGCHLD --> parent
- S (Interruptible Sleep) 可中断睡眠
- I (Idle) 空闲状态
- T (Traced) 暂停/跟踪状态
    SIGSTOP -> Stopped
    SIGCONT -> Runnable
- X (Dead) 消亡，不会在 top 中看到

### 2. 案例分析
> 准备：
> 机器：Ubuntu 18.04, 2 CPU, 8GB
> 工具：apt install docker.io sysstat
> - docker
> - dstat: 吸引 vmstat, iostat, ifstat 等工具优点，可同时观察 CPU、磁盘 IO、网络、内存使用情况
> - systat
> - perf

安装应用实例
```bash
$ docker run --privileged --name=app -itd feisky/app:iowait
```
查看案例是否正常启动
```bash
# s 表示会话领导进程，+ 表示前台进程组
$ ps aux | grep /app
root      2832  0.0  0.0   4512  1556 pts/0    Ss+  12:14   0:00 /app
root      2896  0.0  0.8  70052 65824 pts/0    D+   12:15   0:00 /app
root      2897  0.0  0.8  70052 65824 pts/0    D+   12:15   0:00 /app
```