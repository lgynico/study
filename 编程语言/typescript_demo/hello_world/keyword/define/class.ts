// class 定义一个类

class AClass {
    a: string // 普通属性
    b?: number // 可选属性
    c = true // 推导属性
    d!: symbol // 非空属性
    readonly e: string = "foo" // 只读属性
    static f: string // 静态成员
    #g: number // 私有成员

    // 构造函数
    constructor() {
        this.e = "bar"
    }

    // 普通函数
    add(x: number) {
        return this.b || 0 + x
    }
}


// 构造函数重载
class BClass {
    constructor(x: number, y: string)
    constructor(x: string)
    constructor(x: number | string, y?: string) {
    }
}


// 存储器方法
class CClass {
    _name = ""
    get name() { return this._name }
    // set name(value) { this._name = value }
    set name(value: number | string) { this._name = String(value) }
    // ERROR: Type 'number' is not assignable to type 'string'.
    // set name(value: number) { this._name = value }
}


// 属性索引
class DClass {
    [s: string]: boolean | ((s: string) => boolean)
    get(s: string) { return this[s] as boolean }
}


interface IClassE { a: number, b?: string }
class ClassE implements IClassE { a = 1 }
// ERROR: Class 'ClassE' incorrectly implements interface 'IClassE'.
//   Property 'a' is missing in type 'ClassE' but required in type 'IClassE'.
// class ClassE implements IClassE { }
const a_classE = new ClassE()
a_classE.a
// ERROR: Property 'b' does not exist on type 'ClassE'.
// a_classE.b = "foo"


// 多接口实现，不好
interface IClassF0 { }
interface IClassF1 { }
interface IClassF2 { }
class ClassF implements IClassF0, IClassF1, IClassF2 { }
// 使用接口继承
interface IClassG extends IClassF0, IClassF1, IClassF2 { }
class ClassG implements IClassG { }


interface ClassH { a: number }
class ClassH { b: number = 1 }
const a_classH = new ClassH()
a_classH.a = 10
a_classH.b = 20


//
class ClassI {
    constructor(public x: number, public y: number) { }
}
// function createClassI(ClzClassI: ClassI, x: number, y: number) {
// ERROR: This expression is not constructable.
//   Type 'ClassI' has no construct signatures.
// return new ClzClassI(x, y)
// }
function createClassI(ClzClassI: typeof ClassI, x: number, y: number) {
    return new ClzClassI(x, y)
}


// 结构类型原则
class ClassJ {
    id!: number
}
const b_class = { id: 10, name: "nico" }
const c_class: ClassJ = b_class
console.log(b_class instanceof ClassJ) // false



class ClassKSuper {
    protected x: string = ""
    protected y: string = ""
    protected z: string = ""
    greet() { console.log("hello world") }
}
class ClassK extends ClassKSuper {
    // ERROR: Class 'ClassK' incorrectly extends base class 'ClassKSuper'.
    //   Property 'x' is private in type 'ClassK' but not in type 'ClassKSuper'.
    // private x: string = ""
    protected y: string = ""
    public z: string = ""
    greet(name?: string): void {
        if (name === undefined) { super.greet() }
        else { console.log("hello", name) }
    }
    // ERROR: Property 'greet' in type 'ClassK' is not assignable to the same property in base type 'ClassKSuper'.
    //   Type '(name: string) => void' is not assignable to type '() => void'.
    //   Target signature provides too few arguments. Expected 1 or more, but got 0.
    // greet(name: string): void {}
}
const d_class = new ClassK()
d_class.greet()



class ClassL<T> {
    contents: T
    // ERROR: Static members cannot reference class type parameters.
    // static defaultContents: T
    constructor(value: T) { this.contents = value }
}
const e_class = new ClassL("hello")



abstract class ClassMAbs {
    abstract foo: string
    bar: string = ""
    abstract execute(): string
}
class ClassM extends ClassMAbs {
    foo: string;
    execute(): string {
        return "bar"
    }
}



class ClassN {
    name = "ClassN"
    getName(this: ClassN) {
        return this.name
    }
}
const f_class = new ClassN()
const f_class_ret = f_class.getName;
// EEROR: The 'this' context of type 'void' is not assignable to method's 'this' of type 'ClassN'.
// f_class_ret()



class ClassO {
    // ERROR: A 'this' type is available only in a non-static member of a class or interface.
    // static a: this
    isClassA(): this is AClass {
        return this instanceof AClass
    }
}