CREATE TABLE IF NOT EXISTS urls (
    id SERIAL,
    hash varchar(255),
    url varchar(255),
    created_at DATE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS auth (
    id SERIAL,
    uuid varchar(255),
    email varchar(255),
    password varchar(255),
    created_at DATE,
    PRIMARY KEY (id)
);