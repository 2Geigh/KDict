INSERT INTO users (username, password, days_since_last_interaction, days_since_last_meetup)
VALUES ('poopoohead123', '123456', 3, 5)

-- CREATE TABLE IF NOT EXISTS users (
--     id SERIAL PRIMARY KEY, -- serial means it increments upon instantiation, and primary key forces the value to be unique
--     username VARCHAR(255) NOT NULL UNIQUE, -- not null means the value is required, and unique means it cant appear more than once
--     password BYTEA NOT NULL, -- bytea is raw byte data
--     days_since_last_interaction INTEGER NOT NULL DEFAULT 0,
--     days_since_last_meetup INTEGER DEFAULT 0,
--     meetup_plans JSONB,
--     receives_notifications JSONB
-- );