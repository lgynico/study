/* 付钱 */
#include <stdio.h>

int main(void)
{
    printf("Enter a dollar amount: ");
    int dollar;
    scanf("%d", &dollar);

    int bill20 = dollar / 20;
    printf("$20 bills: %d\n", bill20);

    dollar -= 20 * bill20;
    int bill10 = dollar / 10;
    printf("$10 bills: %d\n", bill10);

    dollar -= 10 * bill10;
    int bill5 = dollar / 5;
    printf("$5 bills: %d\n", bill5);

    dollar -= 5 * bill5;
    int bill = dollar / 1;
    printf("$1 bills: %d\n", bill);

    return 0;
}