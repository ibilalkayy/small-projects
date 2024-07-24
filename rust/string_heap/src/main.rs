fn main() {
    let mut _content: String = String::from("This is the text content.");
    let added_content: String = append(_content.clone());
    insert(added_content.clone(), " This is the inserted text");
    remove(added_content);
}

fn append(mut _content: String) -> String {
    let _add: &str = " This is the appended text.";
    _content.push_str(_add);
    println!("Appended content: {}", _content);
    return _content;
}

fn insert(mut _data: String, _insert: &str) {
    _data.replace_range(0..1, _insert);
    println!("Inserted content: {}", _data);
}

fn remove(mut _content: String) {
    if !_content.is_empty() {
        _content.truncate(_content.len() -1);
        println!("Removed content: {:?}", _content);
    } else {
        println!("The content is empty. Nothing to remove!");
    }
}

