// 数组

let a_array: number[] = [1, 2, 3];
let b_array: (number | string)[] = [1, "2"];
let c_array: Array<number> = [1, 2, 3];
let d_array: Array<number | string> = [1, "2"];


// 成员动态变化
let e_array: number[];
e_array = [];
e_array = [1];
e_array = [1, 2];
e_array = [1, 2, 3];
e_array[5] = 5;
e_array.length = 2;
console.log(e_array);

let e_array_f = e_array[99];


// 读取成员类型
type T_Array = string[];
type T2_Array = T_Array[0]; // string
type T3_Array = T_Array[number]; // string


// 类型推断
let f_array = [];
console.log(typeof f_array, f_array); // any[]
f_array.push(123);
console.log(typeof f_array, f_array); // number[]
f_array.push("abc");
console.log(typeof f_array, f_array); // (number | string)[]


// 不类型推断
let g_array = [123];
g_array.push(456);
// error: Argument of type 'string' is not assignable to parameter of type 'number'.
// g_array.push("abc");



// 只读数组
let h_array: readonly number[] = [0, 1];
// error: Index signature in type 'readonly number[]' only permits reading.
// h_array[1] = 2;
// delete h_array[0];

// error: Property 'push' does not exist on type 'readonly number[]'.
// h_array.push(3);



// readonly 是父类型
let i_array: number[] = [0, 1];
let j_array: readonly number[] = i_array;
// error: The type 'readonly number[]' is 'readonly' and cannot be assigned to the mutable type 'number[]'.
// i_array = j_array;

function arrayTest(a: number[]) { }
// error: Argument of type 'readonly number[]' is not assignable to parameter of type 'number[]'.
//   The type 'readonly number[]' is 'readonly' and cannot be assigned to the mutable type 'number[]'.
// arrayTest(j_array);


// 泛型只读数组
// error: 'readonly' type modifier is only permitted on array and tuple literal types.
// let k_array: readonly Array<number> = [0, 1];
let k_array: ReadonlyArray<number> = [0, 1];
let l_array: Readonly<number[]> = [0, 1];
let m_array = [0, 1] as const;



// 多维数组
let n_array: number[][] = [[1, 2, 3], [4, 5, 6], [7, 8, 9]];