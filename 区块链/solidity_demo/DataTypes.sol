// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract DataTypes {
    int i = 1;
    int8 i8 = 1;
    int16 i16 = 1;
    int32 i32 = 1;
    int64 i64 = 1;
    int128 i128 = 1;
    int256 i256 = 1;

    uint u = 1;
    uint8 u8 = 1;
    uint16 u16 = 1;
    uint32 u32 = 1;
    uint64 u64 = 1;
    uint128 u128 = 1;
    uint256 u256 = 1;

    address addr = 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;

    bytes1 b1 = 0xFF;
    bytes2 b2 = 0xFFFF;
    // ...
    bytes4 b4 = 0xffffffff;
    // ...
    bytes8 b8 = 0xffffffffffffffff;
    // ...
    bytes16 b16 = 0xffffffffffffffffffffffffffffffff;
    // ...
    bytes31 b31 =
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;
    bytes32 b32 =
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;

    bytes b = "hello world";
}
