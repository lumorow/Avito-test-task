CREATE TABLE IF NOT EXISTS segments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_segment_relationship (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    segment_id INTEGER REFERENCES segments(id),
    UNIQUE (user_id, segment_id)
);
