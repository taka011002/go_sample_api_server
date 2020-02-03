CREATE TABLE characters (
    id INTEGER auto_increment PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    character_rarity_id INTEGER NOT NULL
);