CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    balance FLOAT DEFAULT 0 
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    bill FLOAT NOT NULL,
    description VARCHAR(255) NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);