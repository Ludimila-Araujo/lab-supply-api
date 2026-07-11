CREATE TABLE
IF NOT EXISTS order_items
(
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    unit_price NUMERIC
(10,2) NOT NULL,

    PRIMARY KEY
(order_id, product_id),

    CONSTRAINT fk_item_order
        FOREIGN KEY
(order_id)
        REFERENCES orders
(id),

    CONSTRAINT fk_item_product
        FOREIGN KEY
(product_id)
        REFERENCES products
(id)
);