package postgres

import (
	"errors"
	"fmt"
)

func (db *Repo) CreateSegment(segmentName string) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (name) values $1 RETURNING id", "segments")
	var segmentID int

	row, err := db.Db.Query(createSegment, segmentName)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&segmentID)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("Failed to create segment")
	}

	return segmentID, nil
}

func (db *Repo) DeleteSegment(segmentName string) (int, error) {
	deleteSegment := fmt.Sprintf("DELETE FROM %s WHERE name = $1 RETURNING id", "segments")
	var segmentID int

	row, err := db.Db.Query(deleteSegment, segmentName)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&segmentID)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("Failed to delete segment")
	}

	return segmentID, nil
}
