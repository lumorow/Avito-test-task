package postgres

import (
	"Avito-test-task/config"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	cfgDb := cfg.PostgresDB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfgDb.Host, cfgDb.Port, cfgDb.User, cfgDb.Password, cfgDb.Dbname, cfgDb.SSLMode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("ping DB successfully")
	return db, nil
}
