CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    password text NOT NULL,
    email_verified_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);
