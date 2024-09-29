// 泛型

function gen_func_a<T>(value: T) { }
gen_func_a("123")
gen_func_a(123)
gen_func_a(true)


function gen_func_b<T>(arg: T): T { return arg }
const gen_var_a: <T>(arg: T) => T = gen_func_b
const gen_var_b: { <T>(arg: T): T } = gen_func_b



interface gen_interface<T> {
    compareTo(value: T): number
}
class gen_interface_impl implements gen_interface<string> {
    compareTo(value: string): number {
        return 1
    }
}


interface gen_interface_b {
    <T>(arg: T): T
}


const gen_constructor = class <T> {
    constructor(private readonly data: T) { }
}
const gen_constructor_a = new gen_constructor<string>("123")
const gen_constructor_b = new gen_constructor<boolean>(true)



function gen_func_c<T = string>(arr: T[]): T { return arr[0] }

// ERROR: Required type parameters may not follow optional type parameters.
// class gen_class<T = number, U> { }
class gen_class<T, U = number> {
    a: T
    b: U
}


type gen_type<T> = T | null | undefined


class gen_class_b<T extends { length: number }> {
    compare(a: T, b: T): boolean {
        return a.length > b.length
    }
}