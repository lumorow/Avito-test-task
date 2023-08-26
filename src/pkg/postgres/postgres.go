package postgres

import (
	"Avito-test-task/config"
	"database/sql"
	"fmt"
	"io/ioutil"
)

type Repo struct {
	Db *sql.DB
}

func NewPostgresDB(cfg *config.Config) (*Repo, error) {
	R := &Repo{}
	cfgDb := cfg.PostgresDB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfgDb.Host, cfgDb.Port, cfgDb.User, cfgDb.Password, cfgDb.Dbname, cfgDb.SSLMode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	R.Db = db

	query, err := ioutil.ReadFile("./schema/up.sql")
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(string(query)); err != nil {
		return nil, err
	}

	return R, nil
}
