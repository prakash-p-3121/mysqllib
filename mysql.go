package mysqllib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prakash-p-3121/tomllib"
	"io/ioutil"
	"log"
	"time"
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

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.UserName,
		cfg.Password,
		cfg.HostAddr,
		cfg.Port,
		cfg.DatabaseName)
	log.Println("connectionStr=" + connectionString)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func CreateDatabaseConnectionWithRetry(cfgPath string) (*sql.DB, error) {
	var err error
	var db *sql.DB
	for i := 1; i <= 10; i++ {
		db, err = CreateDatabaseConnection(cfgPath)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		return db, err
	}
	return db, nil
}

func CloseDatabaseConnection(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
