// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/hello/print.html
use chrono::{Local, Utc};

#[derive(Debug)]
struct TestStructure(i32);

fn main() {
    println!("Hello, {name}!", name="Benton");
    println!("{:?}", Utc::now());
    println!("{:?}", Local::now());
    println!("Today's date in utc: {year:>04}/{day:>02}/{month:>02}", year=2020, day=6, month=6);
    println!("{:?}", TestStructure(7));
}
