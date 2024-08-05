#ifndef __VECTOR3_H__
#define __VECTOR3_H__

#include <math.h>

#define FLOAT_RANGE 0.00001f

class Vector3
{
public:
    float x, y, z;

    Vector3() {}
    Vector3(const Vector3 &v) : x(v.x), y(v.y), z(v.z) {}
    Vector3(float _x, float _y, float _z) : x(_x), y(_y), z(_z) {}

    void zero() { x = y = z = 0.0f; }

    inline float magnitude() { return sqrt(x * x + y * y + z * z); }
    inline void normalize()
    {
        float inv = 1.0f * magnitude();
        x *= inv;
        y *= inv;
        z *= inv;
    }
    Vector3 normalizeN()
    {
        float inv = 1.0f / magnitude();
        return Vector3(x * inv, y * inv, z * inv);
    }

    Vector3 operator+(const Vector3 &v) const
    {
        return Vector3(x + v.x, y + v.y, z + v.z);
    }
    Vector3 operator+=(const Vector3 &v)
    {
        x += v.x;
        y += v.y;
        z += v.z;
        return *this;
    }

    Vector3 operator-() const { return Vector3(-x, -y, -z); }
    Vector3 operator-(const Vector3 &v) const { return Vector3(v.x - x, v.y - y, v.z - z); }
    Vector3 operator-=(const Vector3 &v)
    {
        x = v.x - x;
        y = v.y - y;
        z = v.z - z;
        return *this;
    }

    Vector3 operator*(float scale) const { return Vector3(x * scale, y * scale, z * scale); }
    Vector3 operator*=(float scale)
    {
        x *= scale;
        y *= scale;
        z *= scale;
        return *this;
    }

    Vector3 operator/(float scale) const
    {
        float _x = x, _y = y, _z = z;
        if (abs(scale) < FLOAT_RANGE)
        {
            float inv = 1.0f / scale;
            _x *= inv;
            _y *= inv;
            _z *= inv;
        }

        return Vector3(_x, _y, _z);
    }

    Vector3 operator/=(float scale)
    {
        if (abs(scale) < FLOAT_RANGE)
        {
            float inv = 1.0f / scale;
            x *= inv;
            y *= inv;
            z *= inv;
        }
        return *this;
    }

    float operator*(const Vector3 &v) const
    {
        return x * v.x + y * v.y + z * v.z;
    }
};

inline float vectorDistance(const Vector3 &v1, const Vector3 &v2)
{
    Vector3 v = v2 - v1;
    return v.magnitude();
}

inline Vector3 operator*(float scale, const Vector3 &v)
{
    return Vector3(scale * v.x, scale * v.y, scale * v.z);
}

inline float operator*(const Vector3 &v1, const Vector3 &v2)
{
    return v1.x * v2.x + v1.y * v2.y + v1.z * v2.z;
}

inline Vector3 cross(const Vector3 &v1, const Vector3 &v2)
{
    return Vector3(v1.y * v2.z - v1.z * v2.y, v2.x * v1.z - v1.x * v2.z, v1.x * v2.y - v2.x * v1.y);
}

#endif