# Golang

## 基础
### 数据类型
基础类型
```go
var (
    // 整型
    _ int
    _ int8
    _ int16
    _ int32
    _ int64
    _ byte // => int8
    _ rune // => int32

    // 无符号整型
    _ uint
    _ uint8
    _ uint16
    _ uint32
    _ uint64

    // 布尔值
    _ bool = true // false

    // 浮点数
    _ float32
    _ float64

    // 无理数
    _ complex64 = 1 + i
    _ complex128

    // 字符串
    _ string = "hello world"

    // 任何类型
    _ any // => interface{}
)
```
集合
```go
var (
    _ []int
    _ [10]string

    _ map[int]string

    _ chan int
)
```