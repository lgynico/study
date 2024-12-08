fn main() {
    println!("Hello, world!");

    f1();
    f2(5);
    f3(5, 'h');

    let y = {
        let x = 3;
        x + 1
    };
    println!("The value of y is: {y}");

    let five = f4();
    println!("The value of five is: {five}");
}

fn f1() {
    println!("Another function.");
}

fn f2(x: i32) {
    println!("The value of x is: {x}");
}

fn f3(value: i32, unit_lable: char) {
    println!("The measurement is: {value}{unit_lable}")
}

fn f4() -> i32 {
    5
}
