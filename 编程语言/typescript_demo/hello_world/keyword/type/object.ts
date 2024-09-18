// object 对象类型，包含对象、数组、函数

let a_object: object;
a_object = { foo: 123 };
a_object = [1, 2, 4];
a_object = (n: number) => n + 1;


// error: Type 'xxx' is not assignable to type 'object'.
// a_object = true;
// a_object = 1;
// a_object = "hello";
// a_object = undefined;
// a_object = null;


let b_object: Object;
b_object = true;
b_object = "hello";
b_object = 1;
b_object = { foo: 123 };
b_object = [1, 2];
b_object = (n: number) => n + 1;
// error: Type 'xxx' is not assignable to type 'Object'.
// d_object = null;
// d_object = undefined;

// {} => Object 简写
let c_object: {};


let d_object: Object = { foo: 0 };
let e_object: object = { foo: 0 };
d_object.toString();
e_object.toString();
// error: Property 'foo' does not exist on type 'Object'.
// d_object.foo;
// e_object.foo;