package cmd

import (
	"database/sql"
	"errors"
	"fmt"
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

func generateStruct(table string) (creates string, err error) {
	rows, err := mssql.Query(fmt.Sprintf(`
		select * from information_schema.columns
		where table_schema = '%s'  #表所在数据库
		and table_name = '%s' 
	    order by ordinal_position; #你要查的表
		`, dbname, table))
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

		fs := fmt.Sprintf("%s %s `gorm:\"%s\" json:\"%s\" default:\"%s\"`", lintString(v.ColumnName), ts.TransferType, cn,
			v.ColumnName, cf)
		if v.ColumnComment != "" {
			fs += " // " + v.ColumnComment
		}
		columns = append(columns, fs)
	}
	creates = "package " + tags + "d\n"
	creates += fmt.Sprintf("// {{UpperTableName}} %s struct\ntype {{UpperTableName}} struct{\n%s\n}", table, strings.Join(columns, "\n"))
	return
}

const daotemplate = `
	// TableName sets the insert table name for this struct type
	func (b {{UpperTableName}}) TableName() string {
		return "{{TableName}}"
	}
`

func generateget() string {
	rs := "// {{UpperTableName}}"
	if comment {
		rs += `
	// @tags {{ModuleName}}
	// @Summary 根据id获取{{TableName}}信息
	// @Description 根据id获取{{TableName}}信息
	// @Accept  json
	// @Produce  json
	// @Param id query string true "{{TableName}} id"
	// @Success 200 {object} {{ModuleName}}d.{{UpperTableName}}
	// @Failure 400 {object} ctx.R
	// @Router /api/v1/{{ModuleName}}/{{TableName}} [get]`
	}
	rs += `
		func (u *{{UpModuleName}}) {{UpperTableName}}(c *ctx.Context) {
				var data {{ModuleName}}d.{{UpperTableName}}
	`
	if cache {
		rs += `	
			err := dao.FindByIDCache(c.Query("id"),&data)
	`
	} else {
		rs += `	
			err := dao.FindByID(c.Query("id"),&data)
	`
	}
	rs += `if c.HandlerError(err) {
				return
			}
			c.JSON(200, data)
		}
`
	return rs
}
func generategets() string {
	rs := "// {{UpperTableName}}s"
	if comment {
		rs += `
	// @tags {{ModuleName}}
	// @Summary 获取{{TableName}}列表
	// @Description 获取{{TableName}}列表
	// @Accept  json
	// @Produce  json
	// @Success 200 {array} {{ModuleName}}d.{{UpperTableName}}
	// @Failure 400 {object} ctx.R
	// @Router /api/v1/{{ModuleName}}/{{TableName}}s [get]`
	}
	rs += `
	func (u *{{UpModuleName}}) {{UpperTableName}}s(c *ctx.Context) {
		q := qs.Auto()
		var data []{{ModuleName}}d.{{UpperTableName}}
`
	if cache {
		rs += `
		err := dao.FindsCache(q,&data)
		if c.HandlerError(err) {
			return
		}
		c.JSON(200, gin.H{"data": data, "total": dao.CountCache("{{TableName}}", q)})
	}
`
	} else {
		rs += `
		err := dao.Finds(q,&data)
		if c.HandlerError(err) {
			return
		}
		c.JSON(200, gin.H{"data": data, "total":dao.CountCache("{{TableName}}", q)})
	}
`
	}
	return rs
}

func generatesave() string {
	rs := "// {{UpperTableName}}Save"
	if comment {
		rs += `
	// @tags {{ModuleName}}
	// @Summary 添加或更新{{TableName}}
	// @Description 添加或更新{{TableName}}
	// @Accept  x-www-form-urlencoded
	// @Produce  json
	// @Param document body {{ModuleName}}d.{{UpperTableName}} true "{{TableName}}信息"
	// @Success 400 {object} ctx.R
	// @Failure 400 {object} ctx.R
	// @Router /api/v1/{{ModuleName}}/{{TableName}} [post]`
	}
	rs += `
	func (u *{{UpModuleName}}) {{UpperTableName}}Save(c *ctx.Context) {
		// 获取用户id
		userid := c.GetAdminId()
		var pd {{ModuleName}}d.{{UpperTableName}}
		err := c.UnmarshalFromString(&pd)
		if c.HandlerError(err) {
			return
		}
		// 获取原来的数据
		var opd {{ModuleName}}d.{{UpperTableName}}
		if pd.Id != 0 {
			err = dao.FindByIDCache(pd.Id,&opd)
			if err != nil {
				log.Error(err)
			}
			if pd.Version != opd.Version {
				c.RespError("提交数据已被更新, 请刷新后重试")
				return
			}
		}
		id, err := dao.InsertOrUpdate(&pd)
		if c.HandlerError(err) {
			return
		}
		// 写入操作记录
		_ = admins.InsertOperateRecordSimple(opd, pd, int(id), userid)
	`
	if cache {
		rs += `
			// 清空相关缓存
			dao.ClearCache("{{TableName}}",id)
		`
	}
	rs += `c.JSON(200, ctx.R{Status: 1, Data: id})
	}
`
	return rs
}

func generaterouter() string {
	return `
		// {{TableName}}
		//g.GET("/{{TableName}}", ctx.Handler(rs.{{UpperTableName}}))
		//g.GET("/{{TableName}}s", ctx.Handler(rs.{{UpperTableName}}s))
		//g.POST("/{{TableName}}", ctx.Handler(rs.{{UpperTableName}}Save))
		//g.GET("/{{TableName}}",commons.SetModel(reflect.TypeOf({{ModuleName}}d.{{UpperTableName}}{})), ctx.Handler(commons.Get))
		//g.GET("/{{TableName}}s",commons.SetModel(reflect.TypeOf([]{{ModuleName}}d.{{UpperTableName}}{})), ctx.Handler(commons.Gets))
		//g.POST("/{{TableName}}",commons.SetModel(reflect.TypeOf({{ModuleName}}d.{{UpperTableName}}{})), ctx.Handler(commons.Save))
	`
}

func generatedao(tablename string) (rs string, err error) {
	rs, err = generateStruct(tablename)
	if err != nil {
		return
	}
	rs += daotemplate
	rs = replace(rs, tablename)
	return
}

func generateapi(tablename string) (rs string) {
	rs = "package " + tags + "\n"
	rs += generateget()
	rs += generategets()
	rs += generatesave()
	rs += generaterouter()
	return replace(rs, tablename)
}
func replace(rs, tablename string) string {
	upperModuleName := lintString(tags)
	upperTableName := lintString(tablename)
	rs = strings.Replace(rs, "{{ModuleName}}", tags, -1)
	rs = strings.Replace(rs, "{{UpperModuleName}}", upperModuleName, -1)
	rs = strings.Replace(rs, "{{TableName}}", tablename, -1)
	rs = strings.Replace(rs, "{{UpperTableName}}", upperTableName, -1)
	rs = strings.Replace(rs, "{{UpModuleName}}", strings.ToUpper(tags), -1)
	return rs
}
