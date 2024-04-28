# Task Management System

This is a simple command-line task management system written in Go. Users can add tasks, mark tasks as completed, and view all tasks.

## Features

- **Add Task**: Users can add a new task with a title and description.
- **Mark Task as Completed**: Users can mark a task as completed.
- **View Tasks**: Users can view all tasks along with their details.

## File Structure

The application follows the principles of Clean Architecture, which promotes separation of concerns and modularity. Here's the file structure:

```
task_manager/
    |- entities/
    |   |- task.go
    |- usecases/
    |   |- add_task.go
    |   |- complete_task.go
    |   |- view_tasks.go
    |- interfaces/
    |   |- cli/
    |       |- main.go
```

### Entities

The `entities` package contains the core data structure of the application. In this case, it defines the `Task` entity, which has properties like title, description, and completion status.

### Use Cases

The `usecases` package contains the business logic of the application. Each use case represents a specific action that the user can perform. For example:
- `add_task.go`: Implements the logic for adding a new task.
- `complete_task.go`: Implements the logic for marking a task as completed.
- `view_tasks.go`: Implements the logic for viewing all tasks.

### Interfaces

The `interfaces` package contains the interface adapters, which are responsible for interacting with the outside world. In this case, it provides a command-line interface (CLI) for users to interact with the application.
- `cli`: Contains the main entry point of the application, where users can interact with the task management system through the command line.

## Getting Started

To run the application, follow these steps:

1. Make sure you have Go installed on your machine. If not, you can download it from [here](https://golang.org/dl/).

2. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/ibilalkayy/small-projects.git
    ```

3. Navigate to the project directory:

    ```bash
    cd task-clean
    ```

4. Run the application:

    ```bash
    go run main.go
    ```