// number 数值类型，包含整数与浮点数

let x_number: number = 123;
let y_number: number = 3.14;
let z_number: number = 0xffff;
let a_number: number = 0b00011;
let b_number: number = 0o721;

// object
let c_number = new Number(123);
// number
let d_number = Number(123);
console.log(typeof c_number, typeof d_number);



let e_number: number = 1;
let f_number: Number = 1;
Math.abs(e_number);
// error: Argument of type 'Number' is not assignable to parameter of type 'number'.
//   'number' is a primitive, but 'Number' is a wrapper object. Prefer using 'number' when possible.
// Math.abs(f_number);

