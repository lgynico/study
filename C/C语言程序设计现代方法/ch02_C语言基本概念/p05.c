/* 求多项式 */
#include <stdio.h>

int main(void)
{
    int x;
    printf("Enter an integer: ");
    scanf("%d", &x);

    int x2 = x * x;
    int x3 = x2 * x;
    int x4 = x3 * x;
    int x5 = x4 * x;
    int result = 3 * x5 + 2 * x4 - 5 * x3 - x2 + 7 * x - 6;

    printf("Result is: %d\n", result);

    return 0;
}