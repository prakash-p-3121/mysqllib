package mysqllib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	database_clustermgt_client "github.com/prakash-p-3121/database-clustermgt-client"
	model "github.com/prakash-p-3121/database-clustermgt-model"
	"github.com/prakash-p-3121/tomllib"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

func getMySQLCfg(filePath string) (*MySQLCfg, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	log.Println("cfg=" + string(data))
	var cfg MySQLCfg
	err = tomllib.Serialize(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func CreateDatabaseConnectionByShard(cfg *model.DatabaseShard) (*sql.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.UserName,
		cfg.Password,
		cfg.IPAddress,
		cfg.Port,
		cfg.DatabaseName)
	log.Println("connectionStr=" + connectionString)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func CreateDatabaseConnectionByCfg(cfg *MySQLCfg) (*sql.DB, error) {

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

func CreateDatabaseConnectionWithRetryByCfg(cfgPath string) (*sql.DB, error) {
	var err error
	var db *sql.DB
	cfgPtr, err := getMySQLCfg(cfgPath)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		db, err = CreateDatabaseConnectionByCfg(cfgPtr)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		return db, err
	}
	return db, nil
}

func CreateDatabaseConnectionWithRetryByShard(shardPtr *model.DatabaseShard) (*sql.DB, error) {
	var err error
	var db *sql.DB

	for i := 1; i <= 10; i++ {
		db, err = CreateDatabaseConnectionByShard(shardPtr)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(1) * time.Second)
			continue
		}
		return db, err
	}
	return db, nil
}

func CreateShardConnectionsWithRetry(tableList []string) (*sync.Map, error) {
	var shardIDToDatabaseConnectionMap sync.Map
	client := database_clustermgt_client.NewDatabaseClusterMgtClient("127.0.0.1", 3001)
	for _, tableName := range tableList {
		shardPtrList, appErr := client.FindAllShardsByTable(tableName)
		if appErr != nil {
			panic(appErr)
		}
		for _, shardPtr := range shardPtrList {
			db, err := CreateDatabaseConnectionWithRetryByShard(shardPtr)
			if appErr != nil {
				panic(err)
			}
			shardIDToDatabaseConnectionMap.Store(shardPtr.ID, db)
		}
	}
	return &shardIDToDatabaseConnectionMap, nil
}

func CloseDatabaseConnection(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
