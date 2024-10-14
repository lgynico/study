#include <stdio.h>
#include <malloc.h>
#include <stdlib.h>

int main(void)
{
    printf("使用 cat /proc/%d/maps 查看内存分配\n", getpid());

    void *addr = malloc(1);
    printf("内存起始地址：%x\n", addr);
    printf("使用 cat /proc/%d/maps 查看内存分配\n", getpid());

    getchar();

    free(addr);
    printf("释放了地址但 heap 不会释放\n");

    getchar();

    return 0;
}