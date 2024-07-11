# Go CRUD Application with Echo and CSV

This is a simple Go application that demonstrates how to perform CRUD (Create, Read, Update, Delete) operations using the Echo web framework and a CSV file for data storage.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ibilalkayy/small-projects
   cd golang/echo
   ```

2. **Install dependencies:**

   Ensure you have Go installed. Then, run:

   ```sh
   go get -u github.com/labstack/echo/v4
   ```

3. **Create the CSV file:**

   Ensure you have a CSV file named `data.csv` in the same directory as your Go file. If it doesn't exist, the program will create it.

## Usage

1. **Run the application:**

   ```sh
   go run main.go
   ```

2. **The server will be running on `localhost:8080`.**

## API Endpoints

### Create a Record

- **URL:** `/create`
- **Method:** `POST`
- **Payload:**
  ```json
  {
    "name": "John Doe"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe"
  }
  ```

### Read a Record

- **URL:** `/read/:id`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe"
  }
  ```

### Update a Record

- **URL:** `/update/:id`
- **Method:** `PUT`
- **Payload:**
  ```json
  {
    "name": "Jane Roe"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "name": "Jane Roe"
  }
  ```

### Delete a Record

- **URL:** `/delete/:id`
- **Method:** `DELETE`
- **Response:**
  ```json
  {
    "message": "Record deleted"
  }
  ```

## Examples

Here are some example `curl` commands to interact with the API:

### Create a Record

```sh
curl -X POST http://localhost:8080/create -H "Content-Type: application/json" -d '{"name":"John Doe"}'
```

### Read a Record

```sh
curl -X GET http://localhost:8080/read/1
```

### Update a Record

```sh
curl -X PUT http://localhost:8080/update/1 -H "Content-Type: application/json" -d '{"name":"Jane Doe"}'
```

### Delete a Record

```sh
curl -X DELETE http://localhost:8080/delete/1
```

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. For any issues, please open an issue on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](../../LICENSE) file for details.