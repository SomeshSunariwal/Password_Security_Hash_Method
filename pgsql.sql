CREATE TABLE users (
    id serial,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL
)