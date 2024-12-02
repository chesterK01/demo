# Demo
A simple Go application that prints "Hello World" by 2 ways: with and without Docker.
## 1. Run without Docker
### Steps

1. **Set Up the Database**
   ```sh
   Open the MySQL Workbench or terminal
   Run the commands below:
   CREATE DATABASE hello_db;
   USE hello_db;

   Import the database file provided (setup.sql) by running:
   mysql -u root -p hello_db < setup.sql


2. **Start the Application**
   ```sh
   Open a terminal or command prompt in the folder where this application is located.
   Run the following command:
   go run main.go

3. **Access the Application**
   ```sh
   Open your browser and go to:
   http://localhost:8080

## 2. Run with Docker
### Steps

1. **Run the Application**
   ```sh
   Open a terminal or command prompt in the folder where this application is located.
   Run this command to start the application and the database:
   docker-compose up --build
   
   Wait for the system to finish setting everything up.

2. **Access the Application:**
   ```sh
   Open your browser and go to:
   http://localhost:8080

3**Stop the Application**
   ```sh
   docker-compose down

