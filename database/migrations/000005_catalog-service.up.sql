CREATE TYPE SeatType AS ENUM('VIP','FLOOR','')

CREATE TABLE IF NOT EXISTS event_ticket
(
    ticket_id GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    ticket_type SeatType NOT NULL,
    event_id INT NOT NULL
);