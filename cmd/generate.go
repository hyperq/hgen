package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hyperq/hgen/cmd/template"
	"strings"

	"github.com/didi/gendry/scanner"
)

var mssql *sql.DB

type columns struct {
	ColumnName    string `ddb:"COLUMN_NAME"`
	DATETYPE      string `ddb:"DATA_TYPE"`
	ColumnComment string `ddb:"COLUMN_COMMENT"`
	ColumnDefault string `ddb:"COLUMN_DEFAULT"`
	ColumnKey     string `ddb:"COLUMN_KEY"`
}

type tables struct {
	TableComment string `ddb:"Table_COMMENT"`
}

// 生成替换的变量
func generateVar() {
	UpperTableName = lintString(TableName)
	if TagName == TableName {
		OneTableName = "One"
		ListTableName = "List"
		SaveTableName = "Save"
	} else {
		t2 := strings.Trim(strings.Trim(TableName, TagName), "_")
		OneTableName = lintString(t2)
		ListTableName = OneTableName + "s"
		SaveTableName = OneTableName + "Save"
	}
}

func generateStruct() (err error) {
	// 获取 TableColumns
	rows, err := mssql.Query(
		fmt.Sprintf(
			`
		SELECT * FROM information_schema.columns
		WHERE table_schema = '%s'  #表所在数据库
		AND table_name = '%s' 
	    ORDER BY ordinal_position; #你要查的表
		`, dbname, TableName,
		),
	)
	if err != nil {
		return
	}
	var cs []columns
	err = scanner.ScanClose(rows, &cs)
	if err != nil {
		return
	}

	var columns []string
	if len(cs) == 0 {
		err = errors.New("找不到数据")
		return
	}
	for _, v := range cs {
		ts, ok := sqltogotype[v.DATETYPE]
		if !ok {
			fmt.Println(v)
			err = errors.New("暂不支持" + v.DATETYPE)
			return
		}
		var cf = ""
		if v.ColumnDefault != "" {
			cf = v.ColumnDefault
		}
		cn := v.ColumnName
		if v.ColumnKey != "" {
			cn += ";pk"
		}

		fs := fmt.Sprintf(
			"%s %s `gorm:\"%s\" json:\"%s\" default:\"%s\"`", lintString(v.ColumnName), ts.TransferType, cn,
			v.ColumnName, cf,
		)
		if v.ColumnComment != "" {
			fs += " // " + v.ColumnComment
		}
		columns = append(columns, fs)
	}
	TableColumns = strings.Join(columns, "\n")
	// 获取 TableComment
	rows2, err := mssql.Query(
		fmt.Sprintf(
			`
		SELECT * FROM information_schema.tables 
		WHERE table_schema = '%s'  #表所在数据库
		AND table_name = '%s'
		`, dbname, TableName,
		),
	)
	if err != nil {
		return
	}
	var tn tables
	err = scanner.ScanClose(rows2, &tn)
	if err != nil {
		return
	}
	TableComment = tn.TableComment
	return
}

func generateModel() (rs string, err error) {
	err = generateStruct()
	if err != nil {
		return
	}
	rs = replace(template.ModelString)
	return
}

func generateApi() (rs string) {
	return replace(template.ApiString)
}
func replace(rs string) string {
	rs = strings.Replace(rs, "{{TagName}}", TagName, -1)
	rs = strings.Replace(rs, "{{UpperTableName}}", UpperTableName, -1)
	rs = strings.Replace(rs, "{{TableName}}", TableName, -1)
	rs = strings.Replace(rs, "{{OneTableName}}", OneTableName, -1)
	rs = strings.Replace(rs, "{{ListTableName}}", ListTableName, -1)
	rs = strings.Replace(rs, "{{SaveTableName}}", SaveTableName, -1)
	rs = strings.Replace(rs, "{{TableComment}}", TableComment, -1)
	rs = strings.Replace(rs, "{{TableColumns}}", TableColumns, -1)
	rs = strings.Replace(rs, "{{VueModel}}", VueModel, -1)
	return rs
}

func generateAdminApi() (rs string) {
	rs = replace(template.VueApiString)
	return
}

func generateAdminStruct() (rs string, err error) {
	rows, err := mssql.Query(
		fmt.Sprintf(
			`
		select * from information_schema.columns
		where table_schema = '%s'  #表所在数据库
		and table_name = '%s' 
	    order by ordinal_position; #你要查的表
		`, dbname, TableName,
		),
	)
	if err != nil {
		return
	}
	var cs []columns
	err = scanner.ScanClose(rows, &cs)
	if err != nil {
		return
	}

	var columns []string
	if len(cs) == 0 {
		err = errors.New("找不到数据")
		return
	}
	for _, v := range cs {
		ts, ok := sqltotstype[v.DATETYPE]
		if !ok {
			fmt.Println(v)
			err = errors.New("暂不支持" + v.DATETYPE)
			return
		}

		fs := fmt.Sprintf(" %s: %s", v.ColumnName, ts.TransferType)
		if v.ColumnComment != "" {
			fs += " // " + v.ColumnComment
		}
		columns = append(columns, fs)
	}
	VueModel = strings.Join(columns, "\n")
	rs = replace(template.VueModelString)
	return
}
