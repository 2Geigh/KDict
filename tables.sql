-- Create tables in order of dependencies
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY, -- serial means it increments upon instantiation, and primary key forces the value to be unique
    username VARCHAR(255) NOT NULL UNIQUE, -- not null means the value is required, and unique means it cant appear more than once
    password BYTEA NOT NULL, -- bytea is raw byte data
    days_since_last_interaction INTEGER NOT NULL DEFAULT 0,
    days_since_last_meetup INTEGER DEFAULT 0,
    meetup_plans JSONB,
    receives_notifications JSONB
);

CREATE TABLE IF NOT EXISTS friends (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthday TIMESTAMP,
    days_since_last_interaction INTEGER NOT NULL DEFAULT 0,
    days_since_last_meetup INTEGER DEFAULT 0,
    phone_number BYTEA NOT NULL,
    meetup_plans JSONB
);

CREATE TABLE IF NOT EXISTS meetups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    friend_id INTEGER REFERENCES friends(id),
    time TIMESTAMP NOT NULL,
    place VARCHAR(255) NOT NULL,
    friends_attending JSONB
);

-- type User struct {
-- 	id                                         int
-- 	username                                   string
-- 	password                                   hash.Hash64 // To be encrypted
-- 	days_since_you_last_interacted_with_anyone int
-- 	days_since_you_last_hung_out_with_anyone   int
-- 	meetup_plans                               []Meetup
-- 	recievesNotifications                      (map[string]bool)
-- }