// 类型兼容

type T_Compatible = number | string;

// number 是 T_Compatible 的子类
let a_compatible: number = 1;
let b_compatible: T_Compatible = a_compatible;


let c_compatible: "hi" = "hi";
let d_compatible: string = "hello";
d_compatible = c_compatible;
// error: Type 'string' is not assignable to type '"hi"'.
// c_compatible = d_compatible;