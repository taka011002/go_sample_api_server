CREATE TABLE users (
    id INTEGER auto_increment PRIMARY KEY,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    phone VARCHAR(255),
    user_status TINYINT(1)
)