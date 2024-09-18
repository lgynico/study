// type 定义类型别名

type Age = number;
let age: Age = 30;


type Color = "red";
// error: Duplicate identifier 'Color'
// type Color = "blue";
if (Math.random() < 0.5) {
    type Color = "blue";
}


type World = "world";
type Greeting = `hello ${World}`;


// error: Expression expected.
// type T = typeof Date();

// error: 'A_Type' only refers to a type, but is being used as a value here.
// type A_Type = number;
// type B_Type = typeof A_Type;


if (Math.random() < 0.5) {
    type T = number;
    let v: T = 5;
} else {
    type T = string;
    let v: T = "hello";
}