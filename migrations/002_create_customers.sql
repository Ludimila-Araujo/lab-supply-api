CREATE TABLE
IF NOT EXISTS customers
(
    id UUID PRIMARY KEY.
    name VARCHAR
(255) NOT NULL,
    cpf CHAR
(11) UNIQUE NOT NULL,
    birth_date DATE NOT NULL,
    address TEXT NOT NULL,
    email VARCHAR
(255),
    phone VARCHAR
(20) NOT NULL,
    password_hash VARCHAR
(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
);