// const 声明一个常量


const a_const: number = 1;

const b_const = { foo: 1 };

// error: Type 'number' is not assignable to type '5'
// const c_const: 5 = 4 + 1;
const c_count: 5 = (4 + 1) as 5;

// 值类型
const x_const = "https";
const y_const: string = "https";
