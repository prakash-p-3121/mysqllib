package mysqllib

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

const (
	conflictErrNumber = 1062
)

func IsConflictError(err error) bool {
	var mySQLErr *mysql.MySQLError
	ok := errors.As(err, &mySQLErr)
	if !ok {
		return false
	}
	if mySQLErr.Number == conflictErrNumber {
		return true
	}
	return false
}

func RollbackTx(tx *sql.Tx, err error) error {
	if err == nil {
		return nil
	}
	errRollback := tx.Rollback()
	if errRollback != nil {
		return fmt.Errorf("RollbackErr : %w \nCurrentErr 2: %w",
			errRollback, err)
	}
	return err
}
