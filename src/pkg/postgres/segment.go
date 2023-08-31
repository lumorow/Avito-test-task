package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *Repo) CreateSegment(segmentName string) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (id, name) values ($1, $2) RETURNING id", "segments")
	var segmentID int

	row, err := db.Db.Query(createSegment, "(SELECT max(id)+1 FROM users)", segmentName)

	if err != nil {
		return 0, err
	}

	if row.Next() {
		err = row.Scan(&segmentID)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("failed to create segment")
	}

	return segmentID, nil
}

func (db *Repo) DeleteSegment(segmentName string) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return err
	}
	deleteSegmentUserRelation := fmt.Sprintf("DELETE FROM %s AS usr USING segments AS s WHERE s.id = usr.segment_id AND s.name = $1", "user_segment_relationship")

	_, err = tx.Exec(deleteSegmentUserRelation, segmentName)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteSegment := fmt.Sprintf("DELETE FROM %s WHERE name = $1", "segments")

	_, err = tx.Exec(deleteSegment, segmentName)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *Repo) GetIdSegment(segmentName string) (int, error) {
	segmentQuery := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", "segments")

	var segmentID int
	err := db.Db.QueryRow(segmentQuery, segmentName).Scan(&segmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("segment not found")
		}
		return 0, err
	}

	return segmentID, nil
}
