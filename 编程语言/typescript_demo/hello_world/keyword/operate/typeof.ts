// typeof 返回操作数类型的字符串

console.log(typeof undefined);
console.log(typeof null);
console.log(typeof true);
console.log(typeof 1);
console.log(typeof "foo");
console.log(typeof {});
console.log(typeof parseInt);
console.log(typeof Symbol());
console.log(typeof 1n);



const a_typeof = { x: 0 };
// {x: number}
type T0_Type = typeof a_typeof;
// number
type T1_Type = typeof a_typeof.x;


let b_typeof = 1;
let c_typeof: typeof b_typeof; // 值运算，删除
if (typeof b_typeof === "number") { // 类型运算，保留
    c_typeof = b_typeof;
}

