use std::io;

fn add(num1:i32, num2:i32) {
    let result:i32 = num1 + num2;
    println!("Here is the result: {}", result);
}

fn take_input()-> (i32, i32) {
    let mut input_one = String::new();
    let mut input_two = String::new();

    io::stdin().read_line(&mut input_one).expect("failed to read the line");
    io::stdin().read_line(&mut input_two).expect("failed to read the line");

    let input_one:i32 = input_one.trim().parse().expect("please enter a valid number");
    let input_two:i32 = input_two.trim().parse().expect("please enter a valid number");

    return (input_one, input_two);
}

fn main() {
    let (num1, num2) = take_input();
    add(num1, num2);
}
