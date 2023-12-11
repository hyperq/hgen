package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hyperq/hgen/cmd/template"
	"github.com/samber/lo"
)

var TsInterfaces string

func generateTsApi() (rs string, err error) {
	var columns []string
	for _, v := range cs {
		if lo.IndexOf(ignores, v.ColumnName) > -1 {
			continue
		}
		ts, ok := Sql2TsType[v.DATETYPE]
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
	TsInterfaces = strings.Join(columns, "\n")
	rs = replace(template.TsApiString)
	return
}
func generateVueTemp() (rs string) {
	rs = replace(template.VueTemplete)
	return
}

var VueForms string
var VueBasicColumns string

func generateVueModel() (rs string, err error) {
	var columns []string
	var columns2 []string
	for _, v := range cs {
		if lo.IndexOf(ignores, v.ColumnName) > -1 {
			continue
		}
		// form
		formts, ok := Sql2VueFormType[v.DATETYPE]
		if !ok {
			fmt.Println(v)
			err = errors.New("暂不支持" + v.DATETYPE)
			return
		}
		formkeys := formts.TransferType
		if strings.Contains(v.ColumnName, "_id") {
			formkeys = VueFormTransID
		}
		if strings.Contains(v.ColumnName, "_time") {
			formkeys = VueFormTransTime
		}
		formkeys = strings.Replace(formkeys, "{{field}}", v.ColumnName, -1)
		formkeys = strings.Replace(formkeys, "{{label}}", v.ColumnComment, -1)
		columns = append(columns, formkeys)
		//
		tablekeys := `
		{
			title: '{{label}}',
			dataIndex: '{{field}}',
			width: 100,
		},
		`
		if strings.Contains(v.ColumnName, "_time") {
			tablekeys = `
			{
				title: '{{label}}',
				dataIndex: '{{field}}',
				sorter: true,
				format: (text) => {
					return ts2dt(text);
				},
				width: 160,
			},
			`
		}
		tablekeys = strings.Replace(tablekeys, "{{field}}", v.ColumnName, -1)
		tablekeys = strings.Replace(tablekeys, "{{label}}", v.ColumnComment, -1)
		columns2 = append(columns2, tablekeys)
	}
	VueForms = strings.Join(columns, ",")
	VueBasicColumns = strings.Join(columns2, ",")
	rs = replace(template.VueModelString)
	return
}
