#include "Vector3.h"
#include "Matrix3x3.h"
#include <math.h>

inline void Matrix3x3::rotateX(float theta)
{
    float sin, cos;
    sincosf(theta, &sin, &cos);

    _11 = 1.0f, _12 = 0.0f, _13 = 0.0f;
    _21 = 0.0f, _22 = cos, _23 = sin;
    _31 = 0.0f, _32 = -sin, _33 = cos;
}

inline void Matrix3x3::rotateY(float theta)
{
    float sin, cos;
    sincosf(theta, &sin, &cos);

    _11 = cos, _12 = 0.0f, _13 = -sin;
    _21 = 0.0f, _22 = 1.0f, _23 = 0.0f;
    _31 = sin, _32 = 0.0f, _33 = cos;
}

inline void Matrix3x3::rotateZ(float theta)
{
    float sin, cos;
    sincosf(theta, &sin, &cos);

    _11 = cos, _12 = sin, _13 = 0.0f;
    _21 = -sin, _22 = cos, _23 = 0.0f;
    _31 = 0.0f, _32 = 0.0f, _33 = 1.0f;
}

Matrix3x3 operator*(const Matrix3x3 &m1, const Matrix3x3 &m2)
{
    Matrix3x3 m;

    m._11 = m1._11 * m2._11 + m1._12 * m2._21 + m1._13 * m2._31;
    m._12 = m1._11 * m2._21 + m1._12 * m2._22 + m1._13 * m2._23;
    m._13 = m1._11 * m2._31 + m1._12 * m2._32 + m1._13 * m2._33;

    m._21 = m1._21 * m2._11 + m1._22 * m2._21 + m1._23 * m2._31;
    m._22 = m1._21 * m2._21 + m1._22 * m2._22 + m1._23 * m2._23;
    m._23 = m1._21 * m2._31 + m1._22 * m2._32 + m1._23 * m2._33;

    m._31 = m1._31 * m2._11 + m1._32 * m2._21 + m1._33 * m2._31;
    m._32 = m1._31 * m2._21 + m1._32 * m2._22 + m1._33 * m2._23;
    m._33 = m1._31 * m2._31 + m1._32 * m2._32 + m1._33 * m2._33;

    return m;
}

Matrix3x3 &operator*=(Matrix3x3 &m1, const Matrix3x3 &m2)
{
    Matrix3x3 m = m1 * m2;
    return m;
}

Vector3 operator*(const Vector3 &v, const Matrix3x3 &m)
{
    return Vector3(
        v.x * m._11 + v.y * m._21 + v.z * m._31,
        v.x * m._21 + v.y * m._22 + v.z * m._23,
        v.x * m._31 + v.y * m._32 + v.z * m._33);
}

Vector3 &operator*=(Vector3 &v, const Matrix3x3 &m)
{
    Vector3 t = v * m;
    return t;
}