// never 空类型，不包含任何值，不可能的类型

let never_x: never;

function never_fn(x: string | number) {
    if (typeof x === "string") {

    } else if (typeof x === "number") {

    } else {
        // never
        x;
    }
}

function ret_never(): never {
    throw new Error("Error");
}
// bottom type
let v1_never: number = ret_never();
let v2_never: string = ret_never();
let v3_never: boolean = ret_never();