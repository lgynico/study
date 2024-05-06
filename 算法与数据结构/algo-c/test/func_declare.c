#include <stdio.h>

int squre2(int);

int main(void)
{
    double i = 5;
    printf("squre(5) = %d\n", squre(i)); // 没有函数声明，会执行参数提升 output: 1
    printf("squre2(5) = %d\n", squre2(i));

    return 0;
}

int squre(int x)
{
    return x * x;
}

int squre2(int x)
{
    return x * x;
}