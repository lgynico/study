// 交叉类型

// never 不可能同时为 number 和 string
let a_inter: number & string;

// 对象合成
let b_inter: { foo: string } & { bar: string };
b_inter = { foo: "hello", bar: "world" };


// 添加属性
type A_Inter = { foo: number };
type B_Inter = A_Inter & { bar: number };