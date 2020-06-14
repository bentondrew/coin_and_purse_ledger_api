// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/hello/print/print_display/testcase_list.html
use std::fmt;
use chrono::{Utc, Datelike};

#[derive(Debug)]
struct DebugStructure(i32);

struct DisplayStructure(i32);

impl fmt::Display for DisplayStructure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.0)
    }
}

struct ListStructure(Vec<i32>);

impl fmt::Display for ListStructure {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // Get vector out of struct
        let vec = &self.0;
        write!(f, "[")?;
        for (count, v) in vec.iter().enumerate() {
            if count != 0 { write!(f, ", ")?; }
            write!(f, "{}", v)?;
        }
        write!(f, "]")
    }
}

fn main() {
    println!("Hello, {name}!", name="Benton");
    let now = Utc::today();
    println!("Today's date in utc: {year:>04}/{month:>02}/{day:>02}", year=now.year(), month=now.month(), day=now.day());
    println!("Display: {}", DisplayStructure(3));
    println!("Debug: {:?}", DebugStructure(7));
    let vs = ListStructure(vec![1, 2, 3]);
    println!("{}", vs);
}
