use std::io::{self, Write};
use std::collections::HashMap;

struct Item {
    _name: String,
    _quantity: u8,
}

struct Collection {
    _items: HashMap<String, Item>
}

impl Collection {
    fn new() -> Self {
        Collection {
            _items: HashMap::new(),
        }
    }

    fn add_item(&mut self, name: &str, quantity: u8) {
        let item = Item {
            _name: name.to_string(),
            _quantity: quantity,
        };
        
        self._items.insert(name.to_string(), item);
        println!("Added an item: {} and quantity: {}", name, quantity);
    }

    fn update_item(&mut self, name: &str, quantity: u8) {
        if let Some(item) = self._items.get_mut(name) {
            item._quantity = quantity;
            println!("Updated item: {} and quantity: {}", name, quantity);
        } else {
            println!("No item in the collection");
        }
    }

    fn list_item(&self) {
        if self._items.is_empty() {
            println!("There is not item present");
        } else {
            for item in self._items.values() {
                println!("Added item: {} and quantity: {}", item._name, item._quantity);
            }
        }
    }
}

fn enter_choice() -> (String, u8) {
    print!("Enter an item name: ");
    io::stdout().flush().expect("failed to flush the stdout");

    let mut take_name: String = String::new();
    io::stdin().read_line(&mut take_name).expect("failed to read the line");
    let name: &str = take_name.trim();

    print!("Enter your quantity: ");
    io::stdout().flush().expect("failed to flush tghe stdout");

    let mut take_quantity: String = String::new();
    io::stdin().read_line(&mut take_quantity).expect("failed to read the line");
    let quantity: u8 = take_quantity.trim().parse().expect("failed to convert to int");

    return (name.to_string(), quantity);
}

fn main() {
    let mut collection = Collection::new();
    loop {
        println!("1. Add an item");
        println!("2. Update an item");
        println!("3. List an item");
        println!("4. Exit");

        print!("Enter your choice: ");
        io::stdout().flush().expect("failed to flush the stdout");

        let mut take_input: String = String::new();
        io::stdin().read_line(&mut take_input).expect("failed to read the line");
        
        let choice: u8 = take_input.trim().parse().expect("failed to convert to int");

        match choice {
            1=>{
                let (name, quantity) = enter_choice();
                collection.add_item(&name, quantity);
            },
            2=>{
                let (name, quantity) = enter_choice();
                collection.update_item(&name, quantity);
            },
            3=>collection.list_item(),
            4=>break,
            _=>println!("failed to recognize the choice"),
        }
    }
}
