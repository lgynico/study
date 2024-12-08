fn main() {
    f1();
    f2();
}

fn f1() {
    let guess: u32 = "42".parse().expect("Not a number!");
    println!("{guess}");

    let x = 2.0;
    let y: f32 = 3.0;

    println!("{x} {y}");

    let t = true;
    let f: bool = false;

    let c = 'z';
    let z: char = 'Z';
    let heart_eyed_cat = '😻';

    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (t1, t2, y3) = tup;
    let five_hundred = tup.0;
    let six_point_four = tup.1;
    let one = tup.2;

    let a1 = [1, 2, 3, 4, 5];
    let a2: [i32; 5] = [1, 2, 3, 4, 5];
    let a3 = [3; 5];
}

fn f2() {
    let sum = 5 + 10;
    let difference = 95.5 - 4.3;
    let product = 4 * 30;
    let quotient = 56.7 / 32.2;
    let truncated = -5 / 3;
    let remainder = 43 % 5;
}
