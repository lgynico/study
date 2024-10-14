#include <stdio.h>
#include <malloc.h>
#include <string.h>
#include <errno.h>

#define MEM_SIZE 1024 * 1024 * 1024
#define LEN 8

int main(void)
{
    char *addr[LEN];
    int i = 0;
    for (int i = 0; i < LEN; ++i)
    {
        addr[i] = (char *) malloc(MEM_SIZE);
        if (!addr[i])
        {
	    printf("malloc fail: %s\n", strerror(errno));
	    return -1;
        }
	printf("malloc 1g address: 0X%p\n", addr[i]);
    }

    getchar();
    return 0;
}
