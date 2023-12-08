package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewMysql new db
func NewMysql() (msdb *sql.DB, err error) {
	msdb, err = sql.Open("mysql", fmt.Sprintf(Dsn, DbName))
	return
}
