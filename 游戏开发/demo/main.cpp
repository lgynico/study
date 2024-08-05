#include <iostream>
#include "Vector3.h"

using namespace std;

void printVector(const Vector3 &v)
{
    cout << "(" << v.x << ", " << v.y << ", " << v.z << ")" << endl;
}

int main()
{
    Vector3 v1(10, 20, 30);
    printVector(v1);

    Vector3 v2 = -v1;
    printVector(v2);

    v1.zero();
    printVector(v1);

    float mag = v2.magnitude();
    cout << "magitude: " << mag << endl;

    Vector3 v3(10, 20, 30);
    printVector(v3 * 10);
    printVector(10 * v3);
    printVector(v3 / 10);

    Vector3 v4(1, 2, 3);
    printVector(v4 + v3);
    printVector(v4 - v3);

    cout << "distance between v3 and b4: " << vectorDistance(v3, v4) << endl;

    v3.normalize();
    printVector(v3);

    return 0;
}