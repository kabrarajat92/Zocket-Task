CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    product_images TEXT[],
    compressed_images TEXT[]
);
