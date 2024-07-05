use std::io;

fn take_input() -> (i32, i32) {
    let mut input_one:String = String::new();
    let mut input_two:String = String::new();

    io::stdin().read_line(&mut input_one).expect("failed to read the line");
    io::stdin().read_line(&mut input_two).expect("failed to read the line");

    // convert the string to integer
    let input_one:i32 = input_one.trim().parse().expect("please enter a valid number");
    let input_two:i32 = input_two.trim().parse().expect("please enter a valid number");

    return (input_one, input_two);
}

fn add(num1:i32, num2:i32) {
    let result:i32 = num1 + num2;
    println!("This is the addition result: {}", result);
}

fn subtract(num1:i32, num2:i32) {
    let result:i32 = num1 - num2;
    println!("This is the subtraction result: {}", result);
}

fn multiply(num1:i32, num2:i32) {
    let result:i32 = num1 * num2;
    println!("This is the multiplcation result: {}", result);
}

fn divide(num1:i32, num2:i32) {
    let result:i32 = num1 / num2;
    println!("This is the division result: {}", result);
}

fn main() {
    let (num1, num2) = take_input();
    add(num1, num2);
    subtract(num1, num2);
    multiply(num1, num2);
    divide(num1, num2);
}
