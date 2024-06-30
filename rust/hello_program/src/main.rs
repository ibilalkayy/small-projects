fn main() {
    let arr:[&str; 5] = ["Hello", "World", "Bye", "World", "Hi"];
    write_arr(arr);
    println!("arr={:?}", arr);
}

fn write_arr(mut arr1:[&str; 5]) {
    arr1[0] = "Follow";
    println!("arr1={:?}", arr1);
}
