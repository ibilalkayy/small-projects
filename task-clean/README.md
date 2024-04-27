# Task Management Application

This is a simple task management application built using the Clean Architecture principles. It allows users to add tasks and list all tasks.

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

## File Structure

The application follows a simple file structure:

- **entities**: Contains the definition of the core data structure (`Task`).
- **usecases**: Contains the application-specific business logic or Use Cases (`AddTaskUseCase`, `ListTasksUseCase`).
- **repositories**: Contains the interfaces and concrete implementations for interacting with data (`TaskRepository`, `InMemoryTaskRepository`).
- **main.go**: Entry point of the application.

## Features

- **Add Task**: Users can add tasks to the task list.
- **List Tasks**: Users can view all tasks in the task list.