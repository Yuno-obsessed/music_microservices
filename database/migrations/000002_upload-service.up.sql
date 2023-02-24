CREATE TABLE IF NOT EXISTS uploads
(
    upload_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    upload_name VARCHAR(200) NOT NULL,
    upload_entity VARCHAR(100) NOT NULL,
    FOREIGN KEY user_id
        REFERENCES users (user_id)
        ON DELETE CASCADE
);