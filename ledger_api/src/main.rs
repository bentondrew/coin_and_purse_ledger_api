// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/primitives/tuples.html

use std::fmt;

fn reverse(pair: (i32, bool)) -> (bool, i32) {
    // 'let' can be used to bind the members if a tuple to variables
    let (integer, boolean) = pair;

    (boolean, integer)
}

#[derive(Debug)]
struct Matrix (f32, f32, f32, f32);


impl fmt::Display for Matrix {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        writeln!(f, "{:?}", (self.0, self.1))?;
        write!(f, "{:?}", (self.2, self.3))
    }
}


fn transpose(in_mat: Matrix) -> Matrix {
    Matrix(in_mat.0, in_mat.2, in_mat.1, in_mat.3)
}


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

    // Tuples
    let long_tuple = (1u8, 2u16, 3u32, 4u64,
                      -1i8, -2i16, -3i32, -4i64,
                      0.1f32, 0.2f64, 'a', true);
    println!("long tuple first value: {}", long_tuple.0);
    println!("long tuple second value: {}", long_tuple.1);
    let tuple_of_tuples = ((1u8, 2u16, 2u32), (4u64, -1i8), -2i16);
    println!("tuple of tuples: {:?}", tuple_of_tuples);
    // Default Debug length limit to printing tuples
    // let too_long_tuple = (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13);
    // println!("too long tuple: {:?}", too_long_tuple);
    let pair = (1, true);
    println!("pair is {:?}", pair);
    println!("the reversed pair is {:?}", reverse(pair));
    println!("one element tuple: {:?}", (5u32,));
    println!("just an integer: {:?}", (5u32));
    let example_tuple = (1, "hello", 4.5, true);
    let (a, b, c, d) = example_tuple;
    println!("{:?}, {:?}, {:?}, {:?}", a, b, c, d);
    let matrix = Matrix(1.1, 1.2, 2.1, 2.2);
    println!("{:?}", matrix);
    println!("{}", matrix);
    println!("Matrix:\n{}", matrix);
    println!("Transpose:\n{}", transpose(matrix));
}
