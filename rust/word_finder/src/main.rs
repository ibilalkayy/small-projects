use std::io;

fn data() -> [&'static str; 3] {
    return ["Hello", "World", "Friends"];
}

fn take_input() -> String {
   let mut input_data = String::new();
    io::stdin().read_line(&mut input_data).expect("failed to read the line");
   return input_data.trim().to_string();
}

fn main() {
    let value = take_input();
    let arrays = data();
    let mut found = false;

    for i in arrays {
        if value == i {
            println!("value is found");
            found = true;
            break;
        }
    }

    if !found {
        println!("value not found");
    }
} 
