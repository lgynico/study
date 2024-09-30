// as 类型断言

type T0_As = "a" | "b" | "c"
let a_as = "a"
// ERROR: Type 'string' is not assignable to type 'T0_As'.
// let b_as: T0_As = a_as
let b_as: T0_As = a_as as T0_As
let c_as: T0_As = <T0_As>a_as


const d_as: { x: number } = { x: 0, y: 0 } as { x: number }


const e_as: object = { a: 1, b: 2, c: 3 };
// ERROR: Property 'length' does not exist on type 'object'.
// e_as.length;

// 运行报错
(e_as as Array<string>).length


// 确定 unkonwn 类型
const f_as_unknown: unknown = "Hello world"
// ERROR: Type 'unknown' is not assignable to type 'string'.
// const f_as: string = f_as_unknown
const f_as: string = f_as_unknown as string



// const g_as_union: number | string = "hello"
// const g_as: number = g_as_union as number


const h_as_0 = 1
const h_as: string = h_as_0 as unknown as string


// as const
type T1_As = "JavaScript" | "TypeScript" | "Python"
function as_func(a: T1_As) { }

let i_as_0 = "JavaScript"
// ERROR: Argument of type 'string' is not assignable to parameter of type 'T1_As'.
// as_func(i_as_0)
// ERROR: A 'const' assertions can only be applied to references to enum members, or string, number, boolean, array, or object literals.
// const i_as: T1_As = i_as_0 as const
const i_as_1 = "JavaScript"
as_func(i_as_1)
let i_as_2 = "JavaScript" as const
as_func(i_as_2)
// ERROR: Type '"TypeScript"' is not assignable to type '"JavaScript"'.
// i_as_2 = "TypeScript"


const j_as_0 = { x: 1, y: 1 } // { x: number, y: number }
const j_as_1 = { x: 1 as const, y: 1 } // { x: 1, y: number }
const j_as_2 = { x: 1, y: 1 } as const // { readonly x: 1, readonly y: 1 }



function as_func_b(x: number, y: number) { return x + y }
const k_as_0 = [1, 2]
// ERROR: A spread argument must either have a tuple type or be passed to a rest parameter.
// as_func_b(...k_as_0)
const k_as = [1, 2] as const
as_func_b(...k_as)


enum E_AS { X, Y }
let l_as_0 = E_AS.X // E_AS
let l_as_1 = E_AS.X as const // E_AS.X