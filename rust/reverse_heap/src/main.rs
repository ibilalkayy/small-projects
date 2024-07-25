use std::io;

fn main() {
    make_reverse();
}

fn take_input() -> String {
    let mut input: String = String::new();
    io::stdin().read_line(&mut input).expect("failed to read the line");
    input.trim().to_string()
}

fn make_reverse() {
    let input_data: String = take_input();
    println!("Original string: {}", input_data);

    let reversed: String = input_data.chars().rev().collect();
    println!("Reverse string: {}", reversed);
}
