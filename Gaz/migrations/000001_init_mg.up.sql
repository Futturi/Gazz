CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    birthdate DATE
);

CREATE INDEX idx_users_username ON users(username);

CREATE TABLE subscribers(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    main_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE INDEX idx_subscribers_user_id ON subscribers(user_id);