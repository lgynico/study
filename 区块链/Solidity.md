# Solidity

## 数据类型
### 整型
```solidity
int i = 1;
int8 i8 = 1;
int16 i16 = 1;
int32 i32 = 1;
int64 i64 = 1;
int128 i128 = 1;
int256 i256 = 1;
```
### 无符号整型
```solidity
uint u = 1;
uint8 u8 = 1;
uint16 u16 = 1;
uint32 u32 = 1;
uint64 u64 = 1;
uint128 u128 = 1;
uint256 u256 = 1;
```
### 地址
```solidity
address addr = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
```
### 字节数组
```solidity
bytes1 b1 = 0xFF;
bytes2 b2 = 0xFFFF;
// ...
bytes4 b4 = 0xffffffff;
// ...
bytes8 b8 = 0xffffffffffffffff;
// ...
bytes16 b16 = 0xffffffffffffffffffffffffffffffff;
// ...
bytes31 b31 = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;
bytes32 b32 = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

bytes b = "hello world";
```
## 操作符
### 数学操作符
```solidity
int a = 10;
int b = 5;

int add = a + b;
int sub = a - b;
int mul = a * b;
int div = a / b;

a++;
b--;
```
### 移位操作符
```solidity
int a = 10;

int ls = a << 2;
int rs = a >> 2;
```
### 赋值操作符
```solidity
int a = 10;
int b = 5;

a += b;
a -= b;
a *= b;
a /= b;

a <<= 2;
b >>= 2;
```