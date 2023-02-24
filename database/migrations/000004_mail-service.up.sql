CREATE TABLE IF NOT EXISTS mails
(
    mail_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    recipient VARCHAR(60) NOT NULL,
    subject VARCHAR(200) NOT NULL,
    upload_id BIGINT, -- for qr code?
    date_sent DATE NOT NULL,
    FOREIGN KEY upload_id
        REFERENCES uploads (upload_id)
        ON DELETE CASCADE
);