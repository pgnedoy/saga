CREATE TABLE orders (
    id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    consumer_id VARCHAR(64) NOT NULL,
    quantity int NOT NULL,
    status VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);