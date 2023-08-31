package postgres

import (
	"fmt"
	"os"
	"path/filepath"
)

func (db *Repo) GetUserReport(userUID int, inputTime string) error {
	fileCSV := fmt.Sprintf("tmp/%v.csv", userUID)
	file, err := os.Create(fileCSV)
	defer file.Close()
	if err != nil {
		return err
	}
	absolutePath, _ := filepath.Abs(fileCSV)
	selectUSA := fmt.Sprintf("(SELECT * FROM user_segment_audit usa WHERE usa.user_UID = %d AND usa.stamp < '%s')", userUID, inputTime)
	createReport := fmt.Sprintf("COPY %s TO '%s' WITH CSV DELIMITER '|' HEADER", selectUSA, absolutePath)
	_, err = db.Db.Exec(createReport)
	if err != nil {
		return err
	}
	return nil
}
