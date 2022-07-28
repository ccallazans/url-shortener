CREATE TABLE IF NOT EXISTS urls (
    id SERIAL,
    short VARCHAR(255) UNIQUE NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_by VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    updated_at DATE,
    created_at DATE,
    PRIMARY KEY (id)
);