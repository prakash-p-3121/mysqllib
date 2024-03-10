package mysqllib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prakash-p-3121/tomllib"
	"io/ioutil"
)

func getMySQLCfg(filePath string) (*MySQLCfg, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	var cfg MySQLCfg
	err = tomllib.Serialize(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func CreateDatabaseConnection(cfgPath string) (*sql.DB, error) {
	cfg, err := getMySQLCfg(cfgPath)
	if err != nil {
		panic(err)
	}
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
