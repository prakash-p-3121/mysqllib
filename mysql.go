package mysqllib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateDatabaseConnection(databaseHost, port, databaseName string) (*sql.DB, error) {
	hostPortDatabase := fmt.Sprintf("%s:%s@/%s", databaseHost, port, databaseName)
	log.Println("hostPortDatabase=", hostPortDatabase)
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
