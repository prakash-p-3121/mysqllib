package mysqllib

import (
	"errors"
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
