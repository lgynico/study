// unkown 表示不确定类型，严格的 any

let ux: unknown;
ux = true;
ux = 42;
ux = "Hello World";

// error: Type 'unknown' is not assignable to type 'boolean'.
// let v1: boolean = ux;
// let v2: number = ux;


// error: is of type 'unknown'
// let v1: unknown = { foo: 123 };
// v1.foo;
// let v2: unknown = "hello";
// v2.trim();
// let v3: unknown = (n = 0) => n + 1;
// v3();


let unkonwn_a: unknown = 1;
// unkonwn_a + 1;
unkonwn_a === 1;
if (typeof unkonwn_a === "number") {
    let r = unkonwn_a + 10;
}

let unkonwn_s: unknown = "hello";
// unkonwn_s.length;
if (typeof unkonwn_s === "string") {
    unkonwn_s.length;
}