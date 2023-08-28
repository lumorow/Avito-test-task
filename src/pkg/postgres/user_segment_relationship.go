package postgres

import (
	"database/sql"
	"errors"
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
		id, err := db.CreateSegmentUserRelation(userID, segmentId)
		if id > 0 {
			return errors.New(fmt.Sprintf("segment: '%s' is already owned by the user with uid: %d", segment, userUID))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Repo) CreateSegmentUserRelation(userUID, segmentID int) (int, error) {
	// Проверка на существование сегмента у пользователя
	id, err := db.CheckSegmentUserRelation(userUID, segmentID)
	if id > 0 {
		return id, errors.New("segment already exist")
	}
	if err != nil {
		return 0, err
	}

	createSegmentUserRelation := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) values ($1, $2)", "user_segment_relationship")
	_, err = db.Db.Query(createSegmentUserRelation, userUID, segmentID)
	if err != nil {
		return 0, err
	}

	return 0, nil
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
	defer rows.Close()
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

func (db *Repo) CheckSegmentUserRelation(userUID, segmentID int) (int, error) {
	checkSegmentUserRelation := fmt.Sprintf("SELECT usr.id FROM %s as usr JOIN segments as s ON s.id = usr.segment_id JOIN users as u ON u.id = usr.user_id WHERE u.id = $1 AND s.id = $2", "user_segment_relationship")

	check := 0
	err := db.Db.QueryRow(checkSegmentUserRelation, userUID, segmentID).Scan(&check)
	//defer rows.Close()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("userSegment not found")
		}
		return 0, err
	}
	return check, nil
}
