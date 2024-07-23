use std::io::{self, Write};
use postgres::{Client, NoTls, Error};

fn connect_db () -> Result<(), Error> {
    let client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls);
    client.expect("REASON").batch_execute("
        CREATE TABLE IF NOT EXISTS project (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            description VARCHAR(255) NOT NULL
        );
    ")?;

    println!("Successfully connected to postgreSQL");
    Ok(())
}

fn insert_data() -> Result<(), Error> {
    let mut client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls)?;

    client.execute(
        "INSERT INTO project(name, description) VALUES($1, $2)",
        &[&"Ferris", &"this is the description"],
    )?;
    println!("Data inserted successfully");
    Ok(())
}

fn view_data() -> Result <(), Error> {
    let mut client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls)?;
    let rows  = client.query("SELECT name, description FROM project", &[])?;

    for row in rows {
        let name:String = row.get(0);
        let description:String = row.get(1);

        println!("Name: {}, description: {}", name, description);
    }
   Ok(())
}

fn update_data() -> Result <(), Error> {
    let mut client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls)?;
    let rows = client.query("SELECT name FROM project", &[])?;
    for row in rows {
        let name:String = row.get(0);
        client.execute("UPDATE project SET name=$1, description=$2 WHERE name=$3", &[&"Tim", &"next description", &name])?;
    }
    Ok(())
}

fn delete_data() -> Result <(), Error> {
    let mut client = Client::connect("postgresql://postgres:1122@localhost/workspace", NoTls)?;
    let rows = client.query("SELECT name FROM project", &[])?;
    for row in rows {
        let name:String = row.get(0);
        println!("{}", name);
        client.execute("DELETE FROM project WHERE name=$1", &[&name])?;
        println!("Deleted user named: {}", name);
    }
    Ok(())
}

fn create_project() {
    let _ = connect_db();
    let _ = insert_data();
    println!("project is created");
}

fn view_project() {
    println!("project data is here");
    let _ = view_data();
}

fn update_project() {
    let _ = update_data();
    println!("project is udpated");
}

fn delete_project() {
    let _ = delete_data();
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
