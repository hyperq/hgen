package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperq/hgen/cmd/template"
	"github.com/samber/lo"
)

var GoStructs string

var ignores = []string{"id", "version", "is_delete", "create_time", "modify_time", "create_by", "modify_by"}

func generateGoStruct() (rs string, err error) {
	var columns []string
	for _, v := range cs {
		if lo.IndexOf(ignores, v.ColumnName) > -1 {
			continue
		}
		ts, ok := Sql2GoType[v.DATETYPE]
		if !ok {
			fmt.Println(v)
			err = errors.New("暂不支持" + v.DATETYPE)
			return
		}
		// 默认值
		var df string
		if v.ColumnDefault != "" {
			df = v.ColumnDefault
		}
		// 主键后面拼接pk
		cn := v.ColumnName
		if v.ColumnKey != "" {
			cn += ";pk"
		}

		fs := fmt.Sprintf("%s %s `gorm:\"%s\" json:\"%s\" default:\"%s\"`",
			lintString(v.ColumnName), ts.TransferType, cn, v.ColumnName, df)

		// 字段说明
		if v.ColumnComment != "" {
			fs += " // " + v.ColumnComment
		}
		columns = append(columns, fs)
	}
	GoStructs = strings.Join(columns, "\n")

	rs = replace(template.ModelString)
	return
}

func generateGoApi() (rs string) {
	return replace(template.ApiString)
}
