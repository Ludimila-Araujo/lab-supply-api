CREATE TABLE
IF NOT EXISTS orders
(
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    status VARCHAR
(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_order_customer
        FOREIGN KEY
(customer_id)
        REFERENCES customers
(id)

);