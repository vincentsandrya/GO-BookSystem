-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS Kategori (
    id SERIAL PRIMARY KEY,
    name varchar(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(200),
    modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(200)
);

CREATE TABLE IF NOT EXISTS Buku (
    id SERIAL PRIMARY KEY,
    title varchar(200) NOT NULL,
    description varchar(200),
    image_url varchar(200),
    release_year int, 
    price int, 
    total_page int, 
    thickness varchar(200),
    category_id int REFERENCES Kategori(id),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(200),
    modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(200)
);


CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    username varchar(200),
    password varchar(200),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(200),
    modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_by varchar(200)
);

INSERT INTO Users (username, password, created_by, modified_by) VALUES ('admin', 'password', 'admin', 'admin');

-- +migrate StatementEnd