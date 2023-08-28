CREATE TABLE IF NOT EXISTS segments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    UID INT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS user_segment_relationship (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    segment_id INTEGER REFERENCES segments(id),
    UNIQUE (user_id, segment_id)
);

SELECT 1
FROM user_segment_relationship usr
         JOIN segments s ON s.id = usr.segment_id
         JOIN users u ON u.id = usr.user_id
WHERE u.uid = 2000 AND s.id = 2
LIMIT 1;