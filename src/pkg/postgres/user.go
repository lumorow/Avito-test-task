package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *Repo) CreateUser(userUID int) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", "users")
	var userID int

	row, err := db.Db.Query(createSegment, userUID)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&userID)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("failed to create user")
	}

	return userID, nil
}

func (db *Repo) GetUserId(userUID int) (int, error) {
	segmentQuery := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", "users")

	var segmentID int
	err := db.Db.QueryRow(segmentQuery, userUID).Scan(&segmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return segmentID, nil
}
