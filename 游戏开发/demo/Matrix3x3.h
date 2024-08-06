#ifndef __MATRIX3X3_H__
#define __MATRIX3X3_H__

#include <algorithm>

class Vector3;

class Matrix3x3
{
private:
    float m[9];

public:
    void set(const float _m[9])
    {
        std::copy(_m[0], _m[0] + 9, m[0]);
    }

    float _11, _12, _13;
    float _21, _22, _23;
    float _31, _32, _33;

    void rotateX(float theta);
    void rotateY(float theta);
    void rotateZ(float theta);

    void scale(float sx, float sy, float sz);
    void scale(const Vector3 &v);

    void projectXZ();
    void projectXY();
    void projectYZ();
    void project(const Vector3 &n);

    void reflectXZ();
    void reflectXY();
    void reflectYZ();
    void reflect(const Vector3 &n);

    void shearXY(float s, float t);
    void shearXZ(float s, float t);
    void shearYZ(float s, float t);
};

Matrix3x3 operator*(const Matrix3x3 &m1, const Matrix3x3 &m2);
Matrix3x3 &operator*=(Matrix3x3 &m1, const Matrix3x3 &m2);

Vector3 operator*(const Vector3 &v, const Matrix3x3 &m);
Vector3 &operator*=(Vector3 &v, const Matrix3x3 &m);

#endif