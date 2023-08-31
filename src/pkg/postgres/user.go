package postgres

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *Repo) CreateUser(userUID int) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (id, UID) values ($1, $2) RETURNING id", "users")
	var userID int

	row, err := db.Db.Query(createSegment, "(SELECT max(id)+1 FROM users)", userUID)
	defer row.Close()

	if err != nil {
		return 0, err
	}

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

func (db *Repo) DeleteUser(userUID int) error {
	tx, err := db.Db.Begin()
	if err != nil {
		return err
	}

	deleteSegmentUserRelation := fmt.Sprintf("DELETE FROM %s AS usr USING users AS u WHERE u.id = usr.user_id AND u.UID = $1", "user_segment_relationship")

	_, err = tx.Exec(deleteSegmentUserRelation, userUID)
	if err != nil {
		tx.Rollback()
		return err
	}

	createSegment := fmt.Sprintf("DELETE FROM %s WHERE UID = ($1) RETURNING id", "users")

	_, err = tx.Exec(createSegment, userUID)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (db *Repo) GetUserId(userUID int) (int, error) {
	segmentQuery := fmt.Sprintf("SELECT id FROM %s WHERE UID = $1", "users")

	var segmentID int
	err := db.Db.QueryRow(segmentQuery, userUID).Scan(&segmentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return segmentID, nil
}
