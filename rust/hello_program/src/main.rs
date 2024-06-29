fn main() {
    let mut s1:String = String::from("Hello "); // owner of Hello is s1
    let length = calculate_length(&mut s1); // transfer ownership to s
    println!("This length of {} is {}", s1, length);
}

fn calculate_length(s:&mut String) -> usize { // new owner of the Hello
    s.push_str("World!!");
    let length:usize = s.len();
    return length; // return the length of s to s1
}