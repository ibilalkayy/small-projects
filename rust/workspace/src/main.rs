use std::io::{self, Write};
use postgres::{Client, NoTls, Error};

fn connect_db () {
    let client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls);
    let _ = client.expect("REASON").batch_execute("
        CREATE TABLE IF NOT EXISTS project {
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(255) NOT NULL
        };
    ");
    println!("Successfully connected to postgreSQL");
}

fn insert_data() {
    let client = Client::connect("postgresql://postgres::1122@localhost/workspace", NoTls);

    let _ = client.expect(
        "INSERT INTO project(name, description) VALUES($1, $2)",
        ["Ferris", "this is the description"],
    );
}

fn create_project() {
    println!("project is created");
}

fn view_project() {
    println!("project is viewed");
}

fn update_project() {
    println!("project is udpated");
}

fn delete_project() {
    println!("project is deleted");
}

fn take_input() {
    println!("1. Create a project");
    println!("2. View a project");
    println!("3. Update a project");
    println!("4. Delete a project");

    print!("Enter your choice: ");
    io::stdout().flush().expect("failed to flush stdout");

    let mut take_choice:String = String::new();
    io::stdin().read_line(&mut take_choice).expect("failed to read the line");

    let take_choice:u8 = take_choice.trim().parse().expect("please enter a valid number");
    
    match take_choice {
        1=>create_project(),
        2=>view_project(),
        3=>update_project(),
        4=>delete_project(),
        _=>println!("choice is not recognized"),
    }
}

fn main() {
    take_input();
}
