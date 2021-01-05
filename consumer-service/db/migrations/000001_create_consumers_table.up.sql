CREATE TABLE users (
    id VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    second_name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    status VARCHAR(64) NOT NULL,
    phone VARCHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);