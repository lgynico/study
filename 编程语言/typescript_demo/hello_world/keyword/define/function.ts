// function 声明一个函数

function a_function(s: string): void { console.log("hello " + s); }
function b_function(s: string) { console.log("hello " + s); }
let c_function = function (s: string) { console.log("hello " + s); }
let d_function: (s: string) => void = function (s) { console.log("hello " + s); }


// (string: any, number: any) => number
type T_Funtion = (string, number) => number;


type T1_Function = (s: string) => void;
let e_function: T1_Function = function (s) { console.log("hello " + s); }


// 参数可少不可多
let f_function: (a: number, b: number) => number;
f_function = (a: number) => a;
// ERROR: Type '(a: number, b: number, c: number) => number' is not assignable to type '(a: number, b: number) => number'.
//   Target signature provides too few arguments.Expected 3 or more, but got 2.
// f_function = (a: number, b: number, c: number) => a + b + c;

let g_funciton = (n: number) => 0;
let h_function = (n: number, s: string) => 0;
h_function = g_funciton;
// ERROR: Type '(n: number, s: string) => number' is not assignable to type '(n: number) => number'.
//   Target signature provides too few arguments. Expected 2 or more, but got 1.
// g_funciton = h_function;


function i_function(x: number, y: number) { return x + y; }
let j_function: typeof i_function = function (x, y) { return x + y; };


// 对象写法
let k_function: { (x: number, y: number): number; };
k_function = function (x, y) { return x + y; };


// 函数有属性
function l_function(x: number) { console.log(x); }
l_function.version = "1.0";
let m_function: {
    (x: number): void;
    version: string;
} = l_function;


interface m_function { (a: number, b: number): number; }


// Function 类型
function n_function(f: Function) { return f(1, 2, 3); }


let o_function = (str: string, times: number): string => str.repeat(times);

function p_function(fn: (a: string) => void): void { fn("world"); }

type T2_Function = { name: string };
const q_function = ["apple", "banana", "cat"].map((name): T2_Function => ({ name }));
(name: T2_Function) => ({ name }); // ERROR
name: (T2_Function) => ({ name }); // ERROR


// 可选参数
function r_function(x?: number) { }
r_function();
r_function(1);
r_function(undefined);

function s_function(x: number | undefined) { }
// ERROR: Expected 1 arguments, but got 0.
// s_function();

// ERROR: A required parameter cannot follow an optional parameter.
// let t_function: (a?: number, b: number) => number;
let t_function: (a: number | undefined, b: number) => number;
let u_function: (a: number, b?: number) => number;
u_function = function (x, y) {
    if (y === undefined) { return x; }
    return x + y;
}



// 参数默认值
function v_function(x: number = 0, y: number = 0): [number, number] { return [x, y]; }
v_function();
v_function(1);
v_function(1, 2);
v_function(undefined, 2);

// ERROR: Parameter cannot have question mark and initializer.
// function w_function(x?: number = 0) { }


// 参数解构
function x_function([x, y, z]: [number, number, number]) { console.log(x + y + z); }
function y_function({ x, y, z }: { x: number, y: number, z: number }) { console.log(x + y + z); }

type T3_Function = { a: number; b: number; c: number; };
function z_function({ a, b, c }: T3_Function) { console.log(a + b + c); }


// rest
function ax_function(...nums: number[]) { }
function bx_function(...args: [number, boolean]) { }
function cx_function(...args: [number, boolean?]) { }
function dx_function(n: number, ...m: number[]) { return m.map((x) => n * x); }
function ex_function(...args: [number, ...string[]]) { }
function fx_function(...[str, times]: [string, number]): string { return str.repeat(times); }


// 只读参数
function gx_function(arr: readonly number[]) {
    // ERROR: Index signature in type 'readonly number[]' only permits reading.
    // arr[0] = 0;
}


// 返回 void
function hx_function(): void { console.log("hello"); }
function ix_function() { return; }
// ERROR: Type 'number' is not assignable to type 'void'.
// function jx_function(): void { return 123; }
function kx_function(): void { return undefined; }
// function lx_function(): void { return null; } // not strictNullChecks

type T4_Function = () => void;
let mx_function: T4_Function = () => { return 123; };
// ERROR: The left-hand side of an arithmetic operation must be of type 'any', 'number', 'bigint' or an enum type.
// mx_function() * 2;



// 返回 never
function nx_function(msg: string): never { throw new Error(msg); }
function ox_function(): Error { return new Error("failed"); }
function px_function(): never { while (true) { console.log(123); }; }
function qx_function(): void { console.log("void"); }
// ERROR: A function returning 'never' cannot have a reachable end point.
// function rx_function(): never { console.log("never"); }
function sx_function(x: string | undefined) {
    if (x === undefined) {
        px_function();
    }
    x;
}


// 局部类型
function tx_function(s: string) {
    type message = string;
    let msg: message = "hello " + s;
    return msg;
}
// ERROR: Cannot find name 'message'.
// let tx_function_ret: message = tx_function("world");


// 高阶函数
(n: number) => (m: number) => m * n;


// 重载
function ux_function(str: string): string;
function ux_function(arr: any[]): any[];
function ux_function(strOrArr: string | any[]): string | any[] {
    if (typeof strOrArr === "string") { return strOrArr.split("").reverse().join(); }
    return strOrArr.slice().reverse();
}

class C_Function {
    #data = "";
    add(num: number): this;
    add(bool: boolean): this;
    add(str: string): this;
    add(value: any): this {
        this.#data += String(value);
        return this;
    }
    toString() {
        return this.#data;
    }
}

type T5_Function = {
    (tag: "a"): HTMLAnchorElement;
    (tag: "canvas"): HTMLCanvasElement;
    (tag: "table"): HTMLTableElement;
    (tag: string): HTMLElement;
};

function vx_function(s: string): number;
function vx_function(arr: any[]): number;
function vx_function(x: any): number { return x.length; }

function vx_function_advance(x: any[] | string): number { return x.length; }



// 构造函数
class C1_Function { }
type T6_Function = new () => C1_Function;
function wx_function(c: T6_Function): C1_Function { return new c(); }
let wx_function_ret = wx_function(C1_Function);

type T7_Function = { new(s: string): object; };
// 可作为构造函数或者普通函数
type T8_Function = {
    new(s: string): object;
    (n?: number): number;
};