CREATE TABLE
IF NOT EXISTS products
(
    id UUID PRIMARY KEY,
    name VARCHAR
(255) NOT NULL,
    description TEXT NOT NULL,
    brand VARCHAR
(255) NOT NULL,
    price NUMERIC
(10,2) NOT NULL,
    stock INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);