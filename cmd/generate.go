package cmd

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/didi/gendry/scanner"
)

var SQL *sql.DB
var cs []column

var (
	UpperTableName string
	TableComment   string
)

type column struct {
	ColumnName    string `ddb:"COLUMN_NAME"`
	DATETYPE      string `ddb:"DATA_TYPE"`
	ColumnComment string `ddb:"COLUMN_COMMENT"`
	ColumnDefault string `ddb:"COLUMN_DEFAULT"`
	ColumnKey     string `ddb:"COLUMN_KEY"`
}

type table struct {
	TableComment string `ddb:"Table_COMMENT"`
}

func generate() (err error) {
	rows, err := SQL.Query(
		fmt.Sprintf(
			`
				SELECT * FROM information_schema.columns
				WHERE table_schema = '%s'  #表所在数据库
				AND table_name = '%s' 
				ORDER BY ordinal_position; #你要查的表
			`, DbName, TableName,
		),
	)
	if err != nil {
		return
	}
	err = scanner.ScanClose(rows, &cs)
	if err != nil {
		return
	}

	if len(cs) == 0 {
		err = errors.New("找不到数据")
		return
	}
	rows2, err := SQL.Query(
		fmt.Sprintf(
			`
				SELECT * FROM information_schema.tables 
				WHERE table_schema = '%s'  #表所在数据库
				AND table_name = '%s'
			`, DbName, TableName,
		),
	)
	if err != nil {
		return
	}
	var tn table
	err = scanner.ScanClose(rows2, &tn)
	if err != nil {
		return
	}
	TableComment = tn.TableComment
	return
}
