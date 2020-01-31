CREATE TABLE user_character (
    id INTEGER auto_increment PRIMARY KEY,
    user_id INTEGER NOT NULL,
    character_id INTEGER NOT NULL,
    UNIQUE (user_id, character_id)
);