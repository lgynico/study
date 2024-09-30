// 断言

// 类型断言：见 as


// 非空断言: strictNullChecks
function assert_func(x?: number | null) {
    assert_number(x)
    console.log(x!.toFixed())
}
function assert_number(e?: number | null) {
    if (typeof e !== "number") { throw new Error("Not a number") }
}


class AssertClass {
    // x: number;
    // y: number;
    x!: number;
    y!: number;

    constructor(x: number, y: number) { }
}


// 断言函数：见 asserts