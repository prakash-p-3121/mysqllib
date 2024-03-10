package mysqllib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabaseConnection(cfg MySQLCfg) (*sql.DB, error) {
	hostPortDatabase := fmt.Sprintf("%s:%d@/%s", cfg.DatabaseName, cfg.Port, cfg.DatabaseName)
	db, err := sql.Open("mysql", hostPortDatabase)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func CloseDatabaseConnection(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
