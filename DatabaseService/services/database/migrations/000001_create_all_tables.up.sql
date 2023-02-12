CREATE TABLE IF NOT EXISTS cities
(
    city_id VARCHAR(36) PRIMARY KEY,
    city_name VARCHAR(100) NOT NULL
);
CREATE TABLE IF NOT EXISTS events
(
    event_id VARCHAR(36) PRIMARY KEY,
    band_name VARCHAR(120) NOT NULL,
    event_city_id VARCHAR(36) NOT NULL,
    event_date DATE DEFAULT NULL,
    FOREIGN KEY (event_city_id)
        REFERENCES cities (city_id)
        ON DELETE CASCADE
    );

CREATE TABLE IF NOT EXISTS login
(
    email VARCHAR(60) PRIMARY KEY,
    password VARCHAR(60) NOT NULL
    );

CREATE TABLE IF NOT EXISTS register
(
    register_id VARCHAR(40) PRIMARY KEY,
    username VARCHAR(120) UNIQUE NOT NULL,
    email VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    country VARCHAR(60) NOT NULL
    );