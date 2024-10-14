#include <stdio.h>
#include <malloc.h>

int main(void)
{
    void *addr = malloc(1024 * 1024);
    printf("此 128K 字节的内存起始地址：%x\n", addr);
    printf("使用 cat /proc/%d/maps 查看内存分配\n", getpid());

    getchar();

    free(addr);
    printf("释放了 128K 字节的内存，内存也归还给了操作系统\n");

    getchar();

    return 0;
}