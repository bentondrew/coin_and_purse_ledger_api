// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/hello/print.html
use chrono::{Utc, Datelike};

#[derive(Debug)]
struct TestStructure(i32);

fn main() {
    println!("Hello, {name}!", name="Benton");
    let now = Utc::today();
    println!("Today's date in utc: {year:>04}/{month:>02}/{day:>02}", year=now.year(), month=now.month(), day=now.day());
    println!("{:?}", TestStructure(7));
}
