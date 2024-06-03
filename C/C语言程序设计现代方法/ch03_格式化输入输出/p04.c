#include <stdio.h>

int main(void)
{
    int n1, n2, n3;

    printf("Enter phone number [(xxx) xxx-xxx]: ");
    scanf("(%3d) %3d-%4d", &n1, &n2, &n3);

    printf("You entered %d.%d.%d\n", n1, n2, n3);

    return 0;
}