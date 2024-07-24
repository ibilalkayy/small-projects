fn main() {
    let mut _vector: Vec<i32> = vec![1,2,3,4,5];
    
    _vector.push(6);
    _vector.push(7);
    _vector.push(8);

    println!("The value of the vectors are: {:?}", _vector);

    let _index: &i32 = &_vector[3];
    println!("The vector index value is: {}", _index);

    _vector.pop();

    if _vector.len() > 3 {
        _vector.remove(3);
    } else {
        println!("index out of bounds for vector");
    }

    println!("Updated vector: {:?}", _vector);
}

