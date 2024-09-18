// symbol 表示独一无二的符号类型，类似字符串

let a_symbol: symbol = Symbol();
let b_symbol: symbol = Symbol();
console.log(a_symbol === b_symbol); // false


// ERROR: A variable whose type is a 'unique symbol' type must be 'const'.
// let c_symbol: unique symbol = Symbol();
const c_symbol: unique symbol = Symbol();
const d_symbol = Symbol(); // unique symbol
// ERROR: This comparison appears to be unintentional because the types 'typeof c_symbol' and 'typeof d_symbol' have no overlap.
// console.log(c_symbol === d_symbol);

// 类似于
// const e_symbol: "hello" = "hello";
// const f_symbol: "world" = "world";
// // ERROR: This comparison appears to be unintentional because the types '"hello"' and '"world"' have no overlap.
// // console.log(e_symbol == f_symbol);


// ERROR: Type 'typeof c_symbol' is not assignable to type 'typeof g_symbol'.
// const g_symbol: unique symbol = c_symbol;
const g_symbol: typeof c_symbol = c_symbol;


// 值相等
const h_symbol: unique symbol = Symbol.for("hello");
const i_symbol: unique symbol = Symbol.for("hello");
// console.log(h_symbol == i_symbol);


const j_symbol: unique symbol = Symbol();
let k_symbol: symbol = j_symbol;
// ERROR: Type 'symbol' is not assignable to type 'unique symbol'.
// const l_symbol: unique symbol = k_symbol;


// 属性名只能是 unique symbol
const m_symbol: unique symbol = Symbol();
let n_symbol: symbol = Symbol();
interface ISymbol {
    [m_symbol]: string;
    // ERROR: A computed property name in an interface must refer to an expression whose type is a literal type or a 'unique symbol' type.
    // [n_symbol]: string;
}


class CSymbol {
    // ERROR: A property of a class whose type is a 'unique symbol' type must be both 'static' and 'readonly'.
    // foo: unique symbol = Symbol();
    static readonly bar: unique symbol = Symbol();
}