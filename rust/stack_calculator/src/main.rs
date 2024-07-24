fn main() {
    let _n1: i32 = 34;
    let _n2: i32 = 34;
    let addition_result: i32 = _n1 + _n2;
    let subtraction_result: i32 = _n1 - _n2;
    let multiplication_result: i32 = _n1 * _n2;
    let division_result: i32;
    if _n2 != 0 {
        division_result = _n1 / _n2;
        println!("This is the result of division: {}", division_result);
    } else {
        println!("Division by zero is not allowed.");
    }

    println!("This is the result of addition: {}", addition_result);
    println!("This is the result of subtraction: {}", subtraction_result);
    println!("This is the result of multiplication: {}", multiplication_result);
}
