package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *Repo) CreateSegment(segmentName string) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", "segments")
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

func (db *Repo) GetIdSegment(segmentName string) (int, error) {
	segmentQuery := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", "segments")

	var segmentID int
	err := db.Db.QueryRow(segmentQuery, segmentName).Scan(&segmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("segment not found")
		}
		return 0, err
	}

	return segmentID, nil
}
