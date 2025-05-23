CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price numeric(10, 2) NOT NULL,
    stock INTEGER NOT NULL,
    image_url VARCHAR(255),
    status VARCHAR(20) DEFAULT 'active',
    category_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);