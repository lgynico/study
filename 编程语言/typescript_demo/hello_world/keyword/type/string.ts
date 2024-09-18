// string 字符串类型

let x_string: string = "hello";
let y_string: string = `${x_string} world`;
let z_string: string = 'foo bar';


// 转为 String()
"hello".charAt(1);

// object
let a_string = new String("hello");
// string
let b_string = String("world");
console.log(typeof a_string, typeof b_string);


let c_string: String = "hello";
let d_string: String = new String("hello");
let e_string: string = "world";
// error: Type 'String' is not assignable to type 'string'.
//   'string' is a primitive, but 'String' is a wrapper object. Prefer using 'string' when possible.
// let f_string: string = new String("world");