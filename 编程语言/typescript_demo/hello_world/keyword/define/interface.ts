// interface 定义接口

interface a_interface {
    a: string,
    b?: number,
    readonly c: boolean,
    [prop: number]: string,
}
let a_interface_i: a_interface = { a: "a", b: 1, c: true }

type T0_Interface = a_interface["a"]


interface b_interface {
    [prop: string]: number
    // ERROR: 'number' index type 'string' is not assignable to 'string' index type 'number'.
    // [prop: number]: string
    [prop: number]: number
}


// 方法的不同写法
const c_interface_d = "d"
interface c_interface {
    a(x: boolean): string
    b: (x: boolean) => string
    c: { (x: boolean): string }
    [c_interface_d](x: string): string
}


// 方法重载
interface d_interface {
    f(): number
    f(x: boolean): number
    f(x: string, y: string): string
}


// 独立函数
interface e_interface {
    (x: number, y: number): number
}
let e_interface_a = (x: number, y: number) => x + y;


// 构造函数
interface f_interface {
    new(): f_interface
}


// 继承
interface g_interface_p1 {
    prop: string
}
interface g_interface_p2 {
    attr: number
}
interface g_interface extends g_interface_p1, g_interface_p2 {
    name: string
    // ERROR: Interface 'g_interface' incorrectly extends interface 'g_interface_p2'.
    //   Types of property 'attr' are incompatible.
    //   Type 'string' is not assignable to type 'number'.
    // attr: string
}


type T1_Interface = { name: string, capital: string }
interface h_interface extends T1_Interface { population: number }

class C0_Interface {
    x: string = ""
    y(): boolean { return true }
}
interface i_interface extends C0_Interface { z: number }


// 接口合并（兼容 Javascript）
interface j_interface { x: string }
interface j_interface { y: number }
interface j_interface { z: boolean }
// ERROR: Subsequent property declarations must have the same type.
//   Property 'z' must be of type 'boolean', but here has type 'string'.
// interface j_interface { z: string }


interface k_interface { x(): string }
interface k_interface { x(y: string): string }
interface k_interface { x(y: "hello"): string }
interface k_interface {
    x(a: number): number
    x(a: number, b: number): number
}

interface k_interface_eq {
    x(y: "hello"): string
    x(a: number): number
    x(a: number, b: number): number
    x(y: string): string
    x(): string
}


interface l_interface_a { x: number }
interface l_interface_b { x: string }
declare const l_interface: l_interface_a | l_interface_b
l_interface.x // number | string