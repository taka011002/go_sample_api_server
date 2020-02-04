CREATE TABLE users (
    id INTEGER auto_increment PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    INDEX index_users_on_username (username)
);