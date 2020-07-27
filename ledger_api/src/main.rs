// Rust tutorial
// https://doc.rust-lang.org/rust-by-example/hello/print/fmt.html
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

#[derive(Debug)]
struct City {
    name: &'static str,
    lat: f32,
    lon: f32,
}

impl fmt::Display for City {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let lat_c = if self.lat >= 0.0 {'N'}  else {'S'};
        let lon_c = if self.lat >= 0.0 {'E'}  else {'W'};
        write!(f, "{}: {:.3}°{} {:.3}°{}",
               self.name, self.lat.abs(), lat_c, self.lon.abs(), lon_c)
    }
}

#[derive(Debug)]
struct Color {
    red: u8,
    green: u8,
    blue: u8,
}

impl fmt::Display for Color {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "RGB ({}, {}, {} {:#04X}{:02X}{:02X})",
               self.red, self.green, self.blue, self.red, self.green, self.blue)
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
    for city in [
        City {name: "Dublin", lat: 53.347778, lon: -6.259722},
        City {name: "Oslo", lat: 59.95, lon: 10.75},
        City {name: "Vancouver", lat: 49.25, lon: -123.1},
    ].iter() {
        println!("{}", *city);
    }
    for color in [
        Color {red: 128, green: 255, blue: 90},
        Color {red: 0, green: 3, blue: 254},
        Color {red: 0, green: 0, blue: 0},
    ].iter() {
        println!("{}", *color);
    }
}
