#include <stdio.h>

int main(void)
{
    int itemNumber;
    float unitPrice;
    int y, m, d;

    printf("Enter item number: ");
    scanf("%d", &itemNumber);

    printf("Enter unit price: ");
    scanf("%f", &unitPrice);

    printf("Enter purchase date (mm/dd/yyyy): ");
    scanf("%2d/%2d/%4d", &m, &d, &y);

    printf("Item\tUnit\tPurchase\n");
    printf("\tPrice\tDate\n");
    printf("%d\t$%6.2f\t%2d/%2d/%4d\n", itemNumber, unitPrice, m, d, y);

    return 0;
}