package postgres

import (
	"fmt"
)

func (db *Repo) CreateSegmentsUserRelation(userUID int, segments []string) error {
	userID, err := db.GetUserId(userUID)
	if userID == 0 {
		if userID, err = db.CreateUser(userUID); err != nil {
			return err
		}
	}

	for _, segment := range segments {
		segmentId, _ := db.GetIdSegment(segment)
		err := db.CreateSegmentUserRelation(userID, segmentId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Repo) CreateSegmentUserRelation(userUID, segmentID int) error {

	createSegmentUserRelation := fmt.Sprintf("INSERT INTO %s (id, user_id, segment_id) values ($1, $2, $3)", "user_segment_relationship")
	_, err := db.Db.Query(createSegmentUserRelation, "(SELECT max(id)+1 FROM user_segment_relationship)", userUID, segmentID)
	if err != nil {
		return err
	}

	return nil
}

func (db *Repo) DeleteSegmentsUserRelation(userUID int, segments []string) error {
	userID, err := db.GetUserId(userUID)
	if userID == 0 {
		return err
	}
	for _, segment := range segments {
		segmentId, _ := db.GetIdSegment(segment)
		err = db.DeleteSegmentUserRelation(userID, segmentId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Repo) DeleteSegmentUserRelation(userUID, segmentID int) error {
	deleteSegmentUserRelation := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND segment_id = $2", "user_segment_relationship")

	_, err := db.Db.Query(deleteSegmentUserRelation, userUID, segmentID)
	if err != nil {
		return err
	}

	return nil
}

func (db *Repo) GetUserSegments(userUID int) ([]string, error) {
	deleteSegmentUserRelation := fmt.Sprintf("SELECT s.name FROM %s as usr JOIN segments as s ON s.id = usr.segment_id WHERE usr.user_id = $1", "user_segment_relationship")

	rows, err := db.Db.Query(deleteSegmentUserRelation, userUID)
	if err != nil {
		return nil, err
	}
	segments := make([]string, 0, 5)
	segment := ""
	for rows.Next() {
		err = rows.Scan(&segment)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}
	return segments, nil
}

func (db *Repo) CheckSegmentUserRelation(userUID, segmentID int) (bool, error) {
	checkSegmentUserRelation := `
		SELECT 1
		FROM user_segment_relationship usr
		JOIN segments s ON s.id = usr.segment_id
		JOIN users u ON u.id = usr.user_id
		WHERE u.uid = $1 AND s.id = $2
		LIMIT 1
	`

	var exists bool
	err := db.Db.QueryRow(checkSegmentUserRelation, userUID, segmentID).Scan(&exists)
	if exists == true {
		fmt.Println(userUID, segmentID)
		return exists, nil
	}
	return exists, err
}
