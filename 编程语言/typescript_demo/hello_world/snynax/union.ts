// 联合类型，类型放大

let a_union: string | number;
a_union = 123;
a_union = "123";


let b_union: true | false;
let c_union: "male" | "female";
let d_union: "赤" | "橙" | "黄" | "绿" | "青" | "蓝" | "紫";


let e_union: string | null;
e_union = "nico";
e_union = null;


function unionTest(u: number | string) {
    // error: Property 'toUpperCase' does not exist on type 'string | number'.
    //   Property 'toUpperCase' does not exist on type 'number'.
    // console.log(u.toUpperCase())
    if (typeof u === "string") { // 类型缩小
        console.log(u.toUpperCase())
    } else {
        console.log(u);
    }
}

function unionTest2(scheme: "http" | "https") {
    switch (scheme) {
        case "http":
            return 80;
        case "https":
            return 443;
    }
}