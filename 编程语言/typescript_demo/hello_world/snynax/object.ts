// 对象

let a_object: { x: number; y: number; } = { x: 1, y: 1 };
// ERROR: Property 'z' does not exist on type '{ x: number; y: number; }'.
// a_object.z = 1;
// console.log(a_object.z);
a_object.x = 2;
console.log(a_object.x);
// ERROR: The operand of a 'delete' operator must be optional.
// delete a_object.x;


type T0_Object = { x: number, y: number };
type T1_Object = { x: number; y: number; };
interface I0_Object { x: number; y: number; }


// ERROR: Property 'y' is missing in type '{ x: number; }' but required in type 'T0_Object'.
// let b_object: T0_Object = { x: 1 };
// ERROR: Object literal may only specify known properties, and 'z' does not exist in type 'T0_Object'.
// let b_object: T0_Object = { x: 1, y: 1, z: 1 };
let b_object: T0_Object = { x: 1, y: 1 };

type T2_Object = T0_Object["x"];


interface I1_Object {
    toString(): string;
    prop: number;
}


// 可选属性
type T3_Object = { a: string; b?: string; };
type T4_Object = { a: string; b: string | undefined; };
let c_object: T3_Object = { a: "hello" };
// ERROR: 'c_object.b' is possibly 'undefined'.
// c_object.b.toLowerCase();
c_object.b?.toLocaleLowerCase();
let c_object_x: string = c_object.b ?? "1";


// 只读属性
type T5_Object = {
    readonly x: string;
    readonly y: {
        z: string;
    };
}
let d_object: T5_Object = { x: "hello", y: { z: "world" } };
// ERROR: Cannot assign to 'xxx' because it is a read-only property.
// d_object.x = "hi";
// d_object.y = { z: "nico" };
d_object.y.z = "nico";

type T6_Object = {
    x: string;
    y: {
        z: string;
    };
};
let e_object: T6_Object = d_object;
e_object.x = "hi";
console.log(d_object.x);

let f_object = { name: "nico" } as const;
// ERROR: Cannot assign to 'name' because it is a read-only property.
// f_object.name = "mike";
let g_object: { name: string } = { name: "nico" } as const;
g_object.name = "mike";



// 属性名的索引类型
type T7_Object = { [property: string]: string; };
let h_object: T7_Object = {
    name: "nico", gender: "male",
    // ERROR: Type 'number' is not assignable to type 'string'.
    // age: 30,
    age: "30"
};

type T8_Object = { [property: number]: number; };
let i_object: T8_Object = [1, 2, 3];
let j_object: T8_Object = { 0: 1, 1: 2, 2: 3 };
// ERROR: Property 'length' does not exist on type 'T8_Object'.
// i_object.length;

type T9_Object = {
    [x: string]: string,
    // ERROR: 'number' index type 'boolean' is not assignable to 'string' index type 'string'.
    // [x: number]: boolean,
    // ERROR: Property 'foo' of type 'boolean' is not assignable to 'string' index type 'string'.
    // foo: boolean,
};


// 解构
let k_object = {
    k_object_id: 1,
    k_object_name: "nico",
    k_object_price: 9.9,
};
let { k_object_id, k_object_name, k_object_price } = k_object;
console.log(k_object_id, k_object_name, k_object_price);
// 指定新变量名
let { k_object_id: kobjx, k_object_name: kobjy, k_object_price: kobjz } = k_object;
console.log(kobjx, kobjy, kobjz);



// 结构类型原则
type T10_Object = { x: number; };
type T11_Object = { x: number; y: number; };
let l_object: T10_Object = { x: 1 };
let m_object: T11_Object = { x: 1, y: 2 };
l_object = m_object;
// ERROR: Property 'y' is missing in type 'T10_Object' but required in type 'T11_Object'.
// m_object = l_object;




type T12_Object = {
    title: string;
    darkMode?: boolean;
};
let n_object: T12_Object = {
    title: "MyTitle",
    // ERROR: Object literal may only specify known properties, but 'darkmode' does not exist in type 'T12_Object'.
    darkmode: true
} as T12_Object;


// 最小可选原则
type T13_Object = {
    a?: number;
    b?: number;
    c?: number;
};
let o_object: T13_Object = {
    // ERROR: Object literal may only specify known properties, and 'd' does not exist in type 'T13_Object'.
    // d: 123
};



let p_object_0 = {};
let p_object_1 = { x: 3 };
let p_object_2 = { y: 4 };
let p_object = {
    ...p_object_0,
    ...p_object_1,
    ...p_object_2,
};


interface I2_Object { }
const q_object: I2_Object = { p0: 1, p1: 2 };
// ERROR: Property 'p0' does not exist on type 'I2_Object'.
// q_object.p0;


interface I3_Object { [key: string]: never }
// ERROR: Type 'number' is not assignable to type 'never'.
// const r_object: I3_Object = { p0: 1 };
