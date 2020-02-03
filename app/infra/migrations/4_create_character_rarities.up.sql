CREATE TABLE character_rarities (
    id INTEGER auto_increment PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    rarity INTEGER NOT NULL
);