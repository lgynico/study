/* 显示日期 */
#include <stdio.h>

int main(void)
{
    printf("Enter a date (mm/dd/yyyy): ");

    int m, d, y;
    scanf("%2d/%2d/%4d", &m, &d, &y);

    printf("You entered the date %4d%2d%2d\n", y, m, d);

    return 0;
}