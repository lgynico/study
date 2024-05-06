#include <crtdefs.h>
#include <stdio.h>

void sort(int a[], size_t len)
{
    for (size_t i = len - 1; i > 0; i--)
    {
        for (size_t j = 0; j < i; j++)
        {
            if (a[j] > a[j + 1])
            {
                a[j] = a[j] ^ a[j + 1];
                a[j + 1] = a[j] ^ a[j + 1];
                a[j] = a[j] ^ a[j + 1];
            }
        }
    }
}

int main()
{
    int arr[] = {4, 2, 6, 2, 6, 3, 21, 9, 5, 34, 1, 5};
    size_t len = sizeof(arr) / sizeof(*arr);
    sort(arr, len);
    for (size_t i = 0; i < len; i++)
    {
        printf("%d\t", arr[i]);
    }

    return 0;
}