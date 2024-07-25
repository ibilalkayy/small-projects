use std::collections::HashMap;
use std::io::{self, Write};

struct Item {
    _name: String,
    _quantity: u32,
}

struct Inventory {
    _items: HashMap<String, Item>,
}

impl Inventory {
    fn new() -> Self {
        Inventory {
            _items: HashMap::new(),
        }
    }

    fn add_item(&mut self, name: &str, quantity: u32) {
        let _item = Item {
            _name: name.to_string(),
            _quantity: quantity,
        };
        
        self._items.insert(name.to_string(), _item);
        println!("Added Item name: {} & quantity: {}", name, quantity);
    }

    fn update_item(&mut self, name: &str, quantity: u32) {
        if let Some(item) = self._items.get_mut(name) {
            item._quantity = quantity;
            println!("Updated item: {} with quantity: {}", name, quantity);
        } else {
            println!("Updated item {} not found", name);
        } 
    }

    fn list_items(&self) {
        if self._items.is_empty() {
            println!("No items in inventory");
        } else {
            for item in self._items.values() {
                println!("Item name: {}, Item Quantity: {}", item._name, item._quantity);
            }
        }
    }
}

fn enter_input() -> (String, u32) {
    print!("Enter item name: ");
    io::stdout().flush().expect("Failed to flush stdout");
    
    let mut name = String::new();
    io::stdin().read_line(&mut name).expect("Failed to read line");
    let name = name.trim();

    print!("Enter quantity: ");
    io::stdout().flush().expect("Failed to flush stdout");

    let mut quantity = String::new();
    io::stdin().read_line(&mut quantity).expect("Failed to read line");
    let quantity: u32 = quantity.trim().parse().expect("Please enter a valid number");

    return (name.to_string(), quantity);
}

fn main() {
    let mut inventory = Inventory::new();

    loop {
        println!("\n1. Add an item");
        println!("2. Update an item");
        println!("3. View all items");
        println!("4. Exit");

        print!("Select your choice [1, 2, 3, 4]: ");
        io::stdout().flush().expect("Failed to flush stdout");

        let mut choice = String::new();
        io::stdin().read_line(&mut choice).expect("Failed to read line");
        let choice: u8 = choice.trim().parse().expect("Please enter a valid number");

        match choice {
            1 => {
                let (name, quantity) = enter_input();
                inventory.add_item(&name, quantity);
            },
            2 => {
                let (name, quantity) = enter_input();
                inventory.update_item(&name, quantity);
            },
            3 => inventory.list_items(),
            4 => break,
            _ => println!("Invalid choice. Please select a valid option."),
        }
    }
}
