CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP not NULL, 
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY, 
    name VARCHAR(100) NOT NULL, 
    description TEXT, 
    price NUMERIC(10, 2) NOT NULL, 
    stock_quantity INT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP not NULL, 
    updated_at timestamp,
    deleted_at timestamp
); 
CREATE TABLE usernameser_products (
    id SERIAL PRIMARY KEY, 
    user_id INT references users(id),
    product_id int references products(id)
); 