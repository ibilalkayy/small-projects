# Calculator Application

This is a simple calculator application written in Go, demonstrating the use of Clean Architecture principles to structure the codebase.

## Clean Architecture

Clean Architecture is a software design approach that emphasizes separation of concerns and independence of dependencies. It consists of several layers, each with its own purpose:

1. **Entities**: Represent the core business logic or data of the application.
2. **Use Cases**: Implement the application's business logic, keeping it independent of any external dependencies.
3. **Interface Adapters**: Provide interfaces for interacting with the application, such as user interfaces or external services.
4. **Frameworks & Drivers**: Contain external frameworks, tools, and libraries that the application uses.

This application follows Clean Architecture by organizing its code into separate packages for each layer:

- **Entities**: Defines the core data structure of the application (`entities/operation.go`).
- **Use Cases**: Implements the business logic of the application (`usecases/` directory).
- **Interface Adapters**: Provides a command-line interface for interacting with the application (`main.go`).
- **Frameworks & Drivers**: Utilizes the Go standard library for command-line interface interactions.

## File Structure

```
calculator/
    |- entities/
    |   |- operation.go
    |- usecases/
    |   |- add.go
    |   |- subtract.go
    |   |- multiply.go
    |   |- divide.go
    |- main.go
```

- **entities/**: Contains the definition of the core data structure of the application (`Operation`).
- **usecases/**: Contains the implementation of business logic for various operations (`Add`, `Subtract`, `Multiply`, `Divide`).
- **main.go**: Entry point of the application, providing a command-line interface for user interaction.

## Running the Application

To run the calculator application:

1. Ensure you have Go installed on your system.
2. Clone this repository to your local machine.
3. Navigate to the `calculator` directory in your terminal.
4. Run the application using the following command:

```
go run main.go
```

Follow the on-screen instructions to perform basic arithmetic operations.
