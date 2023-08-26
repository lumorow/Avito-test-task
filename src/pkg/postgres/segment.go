package postgres

import "fmt"

func (db *Repo) CreateSegment(segment string) (int, error) {
	createSegment := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", segment)
	var SegmentId int

	row, err := db.Db.Query(createSegment, segment)
	if err != nil {
		return 0, err
	}

	err = row.Scan(&SegmentId)
	if err != nil {
		return 0, err
	}
	return SegmentId, nil
}
