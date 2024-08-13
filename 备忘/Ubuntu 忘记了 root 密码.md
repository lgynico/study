# Ubuntu 忘记了 root 密码
## Ubuntu 22.04
1. reboot 按住左 shift 键进入 GRUB 菜单
2. 选择 "Advanced options for Ubuntu" -> "xxx(recovery mode)" 进入恢复模式
3. 选择 "root drop to root shell prompt"
4. 执行 ```mount -o remount,rw /``` 挂载文件系统为可写模式
5. 执行 `passwd` 修改密码
6. reboot
