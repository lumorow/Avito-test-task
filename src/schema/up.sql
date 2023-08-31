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

CREATE TABLE IF NOT EXISTS user_segment_audit (
    id SERIAL PRIMARY KEY,
    user_UID INT,
    segment_id VARCHAR(255) NOT NULL,
    operation char(1)   NOT NULL,
    stamp timestamp NOT NULL
);

CREATE OR REPLACE FUNCTION process_user_segment_audit() RETURNS TRIGGER AS $user_segment_audit$
    BEGIN

    IF (TG_OP = 'DELETE') THEN
        INSERT INTO user_segment_audit (id, user_UID, segment_id, operation, stamp) VALUES (
                                                                                            (SELECT max(id)+1 FROM user_segment_audit),
                                              (SELECT u.UID FROM users u JOIN user_segment_relationship usr on u.id = usr.user_id
                                                            JOIN segments s ON s.id = usr.segment_id
                                                            WHERE u.id = OLD.user_id AND s.id = OLD.segment_id),
                                              (SELECT s.name FROM segments s JOIN user_segment_relationship usr on s.id = usr.segment_id
                                                             JOIN users u on u.id = usr.user_id
                                                             WHERE s.id = OLD.segment_id AND s.id = OLD.segment_id),
                                              'D',
                                              now());
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO user_segment_audit (id, user_UID, segment_id, operation, stamp) VALUES (
                                                                                            (SELECT max(id)+1 FROM user_segment_audit),
                                                (SELECT u.UID FROM users u JOIN user_segment_relationship usr on u.id = usr.user_id
                                                            JOIN segments s ON s.id = usr.segment_id
                                                            WHERE u.id = NEW.user_id AND s.id = NEW.segment_id),
                                                (SELECT s.name FROM segments s JOIN user_segment_relationship usr on s.id = usr.segment_id
                                                             JOIN users u on u.id = usr.user_id
                                                             WHERE s.id =NEW.segment_id AND s.id = NEW.segment_id),
                                              'I',
                                              now());
    END IF;
    RETURN NULL;
END
$user_segment_audit$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER user_segment_audit
    AFTER INSERT OR DELETE ON user_segment_relationship
    FOR EACH ROW EXECUTE FUNCTION process_user_segment_audit();

