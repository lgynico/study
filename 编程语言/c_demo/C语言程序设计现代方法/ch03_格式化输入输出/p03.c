#include <stdio.h>

int main(void)
{
    int gs1Prifix, groupIdentifier, publisherCode, itemNumber, checkDigit;

    printf("Enter ISBN: ");
    scanf("%3d-%d-%3d-%5d-%d", &gs1Prifix, &groupIdentifier, &publisherCode, &itemNumber, &checkDigit);

    printf("GS1 prefix: %d\n", gs1Prifix);
    printf("Group identifier: %d\n", groupIdentifier);
    printf("Publisher code: %d\n", publisherCode);
    printf("Item number: %d\n", itemNumber);
    printf("Check digit: %d\n", checkDigit);

    return 0;
}