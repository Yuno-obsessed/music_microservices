CREATE TABLE IF NOT EXISTS login
(
    email VARCHAR(60) PRIMARY KEY,
    password VARCHAR(60) NOT NULL
    );

CREATE TABLE IF NOT EXISTS register
(
    uuid VARCHAR(40) PRIMARY KEY,
    username VARCHAR(120) UNIQUE NOT NULL,
    email VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    country VARCHAR(60) NOT NULL
    );