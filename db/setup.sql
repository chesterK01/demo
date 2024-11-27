CREATE DATABASE demo_db;
USE demo_db;
CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    message VARCHAR(255)
);
INSERT INTO messages (message) VALUES ("Hello World!");