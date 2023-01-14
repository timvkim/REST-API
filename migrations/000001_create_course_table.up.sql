CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    login text NOT NULL,
    password text NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    price int,
    author_id int,

    CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES users (id)
);

