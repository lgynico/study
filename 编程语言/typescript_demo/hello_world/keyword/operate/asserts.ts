// asserts 断言函数

function assert_string(value: unknown) {
    if (typeof value !== 'string') { throw new Error("Not a string") }
}

function assert_string_1(value: unknown): asserts value is string {
    // if (typeof value !== 'number') { throw new Error("Not a number") }
    if (typeof value !== 'string') { throw new Error("Not a string") }
    // ERROR: Type 'boolean' is not assignable to type 'void'.
    // return true
}


function assert_true(x: unknown): asserts x {
    if (!x) { throw new Error(`${x} should be a truthy value.`) }
}