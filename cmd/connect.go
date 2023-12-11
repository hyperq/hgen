package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// NewMysql new db
func NewMysql() (msdb *sql.DB, err error) {
	dsn := fmt.Sprintf(Dsn, DbName)
	msdb, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	return
}
