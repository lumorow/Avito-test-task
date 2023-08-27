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
		err = db.CreateSegmentUserRelation(userID, segmentId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *Repo) CreateSegmentUserRelation(userUID, segmentID int) error {
	createSegmentUserRelation := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) values ($1, $2)", "user_segment_relationship")

	_, err := db.Db.Query(createSegmentUserRelation, userUID, segmentID)
	if err != nil {
		return err
	}

	return nil
}
