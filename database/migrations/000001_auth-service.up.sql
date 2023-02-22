CREATE TYPE Role_Type AS ENUM('customer','artist','band');

CREATE TABLE IF NOT EXISTS roles
(
    role_id SERIAL PRIMARY KEY,
    role_type Role_Type NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens
(
    token_id SERIAL PRIMARY KEY,
    token_value VARCHAR(50) NOT NULL
    );


CREATE TABLE IF NOT EXISTS users
(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(60) NOT NULL,
    password VARCHAR(36) NOT NULL,
    age INT NOT NULL,
    country VARCHAR(100) NOT NULL,
    email VARCHAR(80) NOT NULL,
    role_id INT NOT NULL,
    FOREIGN KEY role_id
    REFERENCES roles (role_id)
    );