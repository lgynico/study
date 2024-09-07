#include <stdio.h>
#include "inplace_swap.h"

void main()
{
    int x = 2;
    int y = 3;
    printf("before: x = %d, y = %d\n", x, y);
    inplace_swap(&x, &y);
    printf("after: x = %d, y = %d\n", x, y);
}