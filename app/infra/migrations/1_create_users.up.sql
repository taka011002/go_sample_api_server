CREATE TABLE users (
    id INTEGER auto_increment PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL ,
    first_name VARCHAR(255) NOT NULL ,
    last_name VARCHAR(255) NOT NULL ,
    email VARCHAR(255) UNIQUE NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    phone VARCHAR(255) NOT NULL ,
    user_status TINYINT(1) NOT NULL DEFAULT 1
);