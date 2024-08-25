# Excel Import and CRUD System

## Description
This is a Golang-based project using the Gin framework to import data from an Excel file, store it in MySQL, and cache the data in Redis. It also provides a simple CRUD interface to manage the data.

## Prerequisites
- Go 1.20 or later
- MySQL
- Redis

## Installation
1. Clone the repository:
    ```sh
    git clone <repository-url>
    cd project-dir
    ```
2. Install dependencies:
    ```sh
    go mod tidy
    ```
3. Set up the `.env` file with your MySQL and Redis configuration.

4. Create the database and table:
    ```sql
    CREATE DATABASE excel_import;
    USE excel_import;

    CREATE TABLE records (
        id INT AUTO_INCREMENT PRIMARY KEY,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        company VARCHAR(255),
        address VARCHAR(255),
        city VARCHAR(255),
        county VARCHAR(255),
        postal VARCHAR(255),
        phone VARCHAR(255),
        email VARCHAR(255),
        web VARCHAR(255)
    );
    ```

## Running the Application
1. Start the application:
    ```sh
    go run cmd/main.go
    ```
2. Use Postman or another API tool to interact with the endpoints:
    - **POST `/upload`**: Upload an Excel file.
    - **GET `/data`**: View imported data.
    - **PUT `/data/:id`**: Update a record.

## License
This project is licensed under the MIT License.
