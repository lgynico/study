#include <stdio.h>

void inplace_swap(int *x, int *y)
{
    *y = *x ^ *y;
    *x = *x ^ *y;
    *y = *x ^ *y;
}

void main()
{
    int x = 2;
    int y = 3;
    printf("before: x = %d, y = %d\n", x, y);
    inplace_swap(&x, &y);
    printf("after: x = %d, y = %d\n", x, y);
}