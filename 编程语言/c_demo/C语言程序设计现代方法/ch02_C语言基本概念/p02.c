/* 计算球体面积 */
#include <stdio.h>
#include <math.h>

#define FACTOR (4.0f / 3.0f)
// #define M_PI 3.1415926f

int main(void)
{
    float radius = 10.0f;
    float area = FACTOR * M_PI * radius * radius * radius;

    printf("ball area is %.2f\n", area);

    return 0;
}