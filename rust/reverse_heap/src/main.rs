use std::io;

fn main() {
    let data: String = make_reverse();
    println!("{}", data);
}

fn take_input() -> String {
    let mut input: String = String::new();
    io::stdin().read_line(&mut input).expect("failed to read the line");
    return input;
}

fn make_reverse() -> String {
    let input_data: String = take_input();
    input_data.chars().rev().collect()
}
