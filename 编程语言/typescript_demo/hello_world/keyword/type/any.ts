// any 表示任意类型

let x: any;
x = 1;
x = true;
x = "foo";

x(1);
x.foo = 100;


function add(x, y) {
    return x + y;
}

add(1, [1, 2, 3]);


let y: number;
y = x;
y * 123;
y.toFixed();