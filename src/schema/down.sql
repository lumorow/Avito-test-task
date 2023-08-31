DROP TABLE IF EXISTS user_segment_audit;
DROP TABLE IF EXISTS user_segment_relationship;
DROP TABLE IF EXISTS segments;
DROP TABLE IF EXISTS users;


DROP TRIGGER IF EXISTS user_segment_audit ON user_segment_relationship;
DROP FUNCTION IF EXISTS process_user_segment_audit();