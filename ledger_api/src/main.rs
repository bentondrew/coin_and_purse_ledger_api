// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/primitives/literals.html

fn main() {
    // Type annotation
    let _logical: bool = true;
    let _a_float: f64 = 1.0; // Regular
    let _an_integer = 5i32; // Suffix (only numbers)
    let _default_float = 3.0; // f64 (only numbers)
    let _default_integer = 7; // i32 (only numbers)
    // Type inference from context
    let mut _inferred_type = 12;
    _inferred_type = 4294967296i64; // This line sets the type at declaration
    // Variable that can be changed
    let mut _mutable = 12; // i32 that can be changed
    _mutable = 21;
    // Variable type cannot be changed
    // _mutable = true;
    // Variables can be overwritten with shadowing
    let _mutable = true; // Redeclares _mutable as bool rather than i32

    // Integer addition
    println!("1 + 2 = {}", 1u32 + 2);
    // Integer subtraction
    println!("1 - 2 = {}", 1i32 - 2);
    // Boolean logic
    println!("true AND false is {}", true && false);
    println!("true OR false is {}", true || false);
    println!("NOT true is {}", !true);
    // Bitwise operations
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u32 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101);
    println!("1 << 5 is {}", 1u32 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2);
    // Underscore in numbers to improve readability
    println!("One million is written as {}", 1_000_000u32);
}
