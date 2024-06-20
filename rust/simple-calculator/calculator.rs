static mut RESULT: i128 = 4;

fn add(num: i128) {
    unsafe {
        RESULT += num;
        println!("{}", RESULT);
    }
}

fn subtract(num: i128) {
    unsafe {
        RESULT -= num;
        println!("{}", RESULT);
    }
}

fn multiple(num: i128) {
    unsafe {
        RESULT *= num;
        println!("{}", RESULT);
    }
}

fn main() {
    add(2);
    subtract(2);
    multiple(2);
}