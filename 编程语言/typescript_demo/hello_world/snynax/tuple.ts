// 元组

let a_tuple: [string, number, boolean] = ["nico", 30, true];

let b_tuple = [1, true]; // array

// 可选成员
let c_tuple: [number, number?] = [1];
type T_Tuple = [number, number, number?, number?];


// 越界
// error: Type '3' is not assignable to type 'undefined'.
// c_tuple[2] = 3;


// 不限成员
type T1_Tuple = [string, ...number[]];
let d_tuple: T1_Tuple = ["a", 1, 3];
let e_tuple: T1_Tuple = ["a", 1, 2, 3, 4];

type T2_Tuple = [string, ...number[], boolean];
type T3_Tuple = [...number[], string, boolean];

// 不推荐
type T4_Tuple = [...any[]];


// 成员类型
type T5_Tuple = [string, number];
type T5_Tuple_E1 = T5_Tuple[1]; // number
type T5_Tuple_E = T5_Tuple[number]; // string | number



// 只读元组
type T6_Tuple = readonly [number, string];
type T7_Tuple = Readonly<[number, string]>;

type T8_Tuple = [number, string];
let f_tuple: T8_Tuple = [1, "A"];
let g_tuple: T7_Tuple = f_tuple;
// error: The type 'readonly [number, string]' is 'readonly' and cannot be assigned to the mutable type 'T8_Tuple'.
// f_tuple = g_tuple;


function testTuple([x, y]: [number, number]) {
    return Math.sqrt(x ** 2 + y ** 2);
}
let h_tuple = [1, 2] as const;
// error: Argument of type 'readonly [1, 2]' is not assignable to parameter of type '[number, number]'.
//   The type 'readonly [1, 2]' is 'readonly' and cannot be assigned to the mutable type '[number, number]'.
// testTuple(h_tuple);
testTuple(h_tuple as [number, number]);



function testTuple2(point: [number, number, number?]) {
    // error: This comparison appears to be unintentional because the types '2 | 3' and '4' have no overlap.
    // if (point.length === 4) { }
}
function testTuple3(t: [...string[]]) {
    if (t.length === 4) {
        console.log(t)
    }
}



function add(x: number, y: number) {
    return x + y
}
let i_tuple = [1, 2];
// error: A spread argument must either have a tuple type or be passed to a rest parameter.
// add(...i_tuple);
console.log(...i_tuple);

let j_tuple: [number, number] = [1, 2];
add(...j_tuple);