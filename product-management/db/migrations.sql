CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_images TEXT[],
    product_price DECIMAL(10, 2),
    compressed_product_images TEXT[]
);
