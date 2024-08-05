#ifndef __MATRIX3X3_H__
#define __MATRIX3X3_H__

class Vector3;

class Matrix3x3
{
public:
    float _11, _12, _13;
    float _21, _22, _23;
    float _31, _32, _33;

    void rotateX(float theta);
    void rotateY(float theta);
    void rotateZ(float theta);
};

Matrix3x3 operator*(const Matrix3x3 &m1, const Matrix3x3 &m2);
Matrix3x3 &operator*=(Matrix3x3 &m1, const Matrix3x3 &m2);

Vector3 operator*(const Vector3 &v, const Matrix3x3 &m);
Vector3 &operator*=(Vector3 &v, const Matrix3x3 &m);

#endif