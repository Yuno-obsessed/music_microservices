CREATE TABLE IF NOT EXISTS event_ticket
(
    event_id INT PRIMARY KEY,
    default_quantity INT,
    vip_quantity INT,
    scene_quantity INT,
    default_cost INT NOT NULL,
    vip_cost INT NOT NULL,
    scene_cost INT NOT NULL,
    FOREIGN KEY event_id
        REFERENCES events (event_id)
        ON DELETE CASCADE
);