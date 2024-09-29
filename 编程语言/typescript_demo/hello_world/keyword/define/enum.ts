// enum 声明一组枚举

enum EnumA {
    EnumA_A,
    EnumA_B,
    EnumA_C = 4,
    EnumA_D = "D",
    EnumA_E = 0.5,
    // ERROR: Type 'bigint' is not assignable to type 'number' as required for computed enum member values.
    // EnumA_F = 7n
    EnumA_G = 1 << 8,
    EnumA_H = Math.random()
}

enum EnumA {
    // ERROR: Duplicate identifier 'EnumA_A'.
    // EnumA_A = 1,
    // ERROR: In an enum with multiple declarations, only one declaration can omit an initializer for its first enum element.
    // EnumA_X,
    EnumA_Z = 1,
}

const enum_a = EnumA.EnumA_A
const enum_b = EnumA["EnumA_B"]
const enum_c: number = EnumA.EnumA_C


// ERROR: Cannot assign to 'EnumA_A' because it is a read-only property.
// EnumA.EnumA_A = 10


// function enum_func(e: EnumA) { }
// enum_func(999)


const EnumB = {
    EnumB_A: 0,
    EnumB_B: 1,
    EnumB_C: 2
} as const
const enum_d = EnumB.EnumB_A



enum EnumC {
    EnumC_A = "a",
    EnumC_B = "b"
}
type EnumC_T1 = typeof EnumC
type EnumC_T2 = keyof EnumC // keyfo number
type EnumC_T3 = keyof typeof EnumC // "EnumC_A" | "EnumC_B"
type EnumC_T4 = { [key in EnumC]: any } // { a：any, b: any }


// 反向映射
enum EnumD {
    Monday = 1,
    Tuesday,
    Wednesday,
    Thursday,
    Friday,
    Saturday,
    Sunday
}
console.log(EnumD[3])