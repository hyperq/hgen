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
	
	const {{UpperTableName}}sql=` + "`" + `
			SELECT a.* 
			FROM {{TableName}} a 
` + "`" + `

	// {{UpperTableName}}s get {{TableName}} list
	func {{UpperTableName}}s(q *qs.QuerySet) (data []{{UpperTableName}}, err error) {
		err = dao.QueryByQs({{UpperTableName}}sql, q,&data)
		return
	}
	// {{UpperTableName}}o get {{TableName}}
	func {{UpperTableName}}o(q *qs.QuerySet) (data {{UpperTableName}}, err error) {
		q.ResetOther()
		datas, err := {{UpperTableName}}s(q)
		if err != nil {
			return
		}
		if len(datas) == 0 {
			err = dao.NotFound
			return
		}
		data = datas[0]
		return
	}
	// {{UpperTableName}}ById get by id
	func {{UpperTableName}}ById(id interface{}) ({{UpperTableName}}, error) {
		q := qs.New()
		q.Add("a.id=?", id)
		return {{UpperTableName}}o(q)
	}
	// {{UpperTableName}}sCache
	func {{UpperTableName}}sCache(q *qs.QuerySet) (data []{{UpperTableName}}, err error) {
		cachekey := q.FormatCache("{{TableName}}l")
		res, err := dao.{{UpperModuleName}}Cache.Get(cachekey, q)
		if err != nil {
			res, err = dao.{{UpperModuleName}}Cache.Add(cachekey, dao.DefaultCacheExpire, func(q *qs.QuerySet) (res string, err error) {
				data, err := {{UpperTableName}}s(q)
				if err != nil {
					return
				}
				res, err = jsoniter.MarshalToString(data)
				return
			}, q)
			if err != nil {
				return
			}
		}
		err = jsoniter.UnmarshalFromString(res, &data)
		return
	}
	// {{UpperTableName}}ByIdCache get cache by id
	func {{UpperTableName}}ByIdCache(id interface{}) (data {{UpperTableName}}, err error) {
		q := qs.New()
		q.Add("a.id=?", id)
		cachekey := "{{TableName}}d" + fmt.Sprint(id)
		res, err := dao.{{UpperModuleName}}Cache.Get(cachekey, q)
		if err != nil {
			res, err = dao.{{UpperModuleName}}Cache.Add(cachekey, dao.DefaultCacheExpire, func(q *qs.QuerySet) (res string, err error) {
				data, err := {{UpperTableName}}o(q)
				if err != nil {
					return
				}
				res, err = jsoniter.MarshalToString(data)
				return
			}, q)
			if err != nil {
				return
			}
		}
		err = jsoniter.UnmarshalFromString(res, &data)
		return
	}
	// {{UpperTableName}}Count count ad number by cache
	func {{UpperTableName}}Count(q *qs.QuerySet) int {
		countkey := q.FormatCache("{{TableName}}c")
		counts, err := dao.{{UpperModuleName}}Cache.Get(countkey, q)
		if err != nil {
			counts, err = dao.{{UpperModuleName}}Cache.Add(countkey, dao.DefaultCacheExpire, func(q *qs.QuerySet) (data string, err error) {
				data = strconv.Itoa(dao.Count("{{TableName}}", q))
				return
			}, q)
			if err != nil {
				return 0
			}
		}
		count, _ := strconv.Atoi(counts)
		return count
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
		func (u *{{UpModuleName}}) {{UpperTableName}}(c *ctx.Context) {`
	if cache {
		rs += `	
			data, err := {{ModuleName}}d.{{UpperTableName}}ByIdCache(c.Query("id"))
	`
	} else {
		rs += `	
			data, err := {{ModuleName}}d.{{UpperTableName}}ById(c.Query("id"))
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
		q := qs.New().Paging(c)
		q.SetArray(c)
		q.SetLikeArray(c)
`
	if cache {
		rs += `
		data, err := {{ModuleName}}d.{{UpperTableName}}sCache(q)
		if c.HandlerError(err) {
			return
		}
		c.JSON(200, gin.H{"data": data, "total": {{ModuleName}}d.{{UpperTableName}}Count(q)})
	}
`
	} else {
		rs += `
		data, err := {{ModuleName}}d.{{UpperTableName}}s(q)
		if c.HandlerError(err) {
			return
		}
		c.JSON(200, gin.H{"data": data, "total": {{ModuleName}}d.Count("{{TableName}}", q)})
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
			opd, err = {{ModuleName}}d.{{UpperTableName}}ById(pd.Id)
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
		pd.Id = int(id)
		// 写入操作记录
		_ = admins.InsertOperateRecordSimple(opd, pd, userid)
	`
	if cache {
		rs += `
			// 清空相关缓存
			dao.{{UpperModuleName}}Cache.FlushIndeof("{{TableName}}l")
			dao.{{UpperModuleName}}Cache.Flush("{{TableName}}d" + strconv.Itoa(pd.Id))
			dao.{{UpperModuleName}}Cache.Flush("{{TableName}}c")
		`
	}
	rs += `c.JSON(200, ctx.R{Status: 1, Data: id})
	}
`
	return rs
}
func generatedelete() string {
	rs := "// {{UpperTableName}}Delete"
	if comment {
		rs += `
	// @tags {{ModuleName}}
	// @Summary 删除数据{{TableName}}
	// @Description 删除数据{{TableName}}
	// @Accept  json
	// @Produce  json
	// @Param id query string true "{{TableName}} id"
	// @Param version query int true  "current version"
    // @Success 400 {object} ctx.R
	// @Failure 400 {object} ctx.R
	// @Router /api/{{ModuleName}}/{{TableName}} [delete]`
	}
	rs += `
	func (u *{{UpModuleName}}) {{UpperTableName}}Delete(c *ctx.Context) {
		ids := strings.Split(c.Query("ids"), ",")
		versions := strings.Split(c.Query("versions"), ",")
		for i := range ids {
			err := dao.Delete("{{TableName}}", ids[i], versions[i])
			if c.HandlerError(err) {
				return
			}
			_ = admins.InsertOperateRecord(2, c.GetAdminId(), "{{TableName}}", ids[i], "")
		}
	`
	if cache {
		rs += `
			// 清空相关缓存
			dao.{{UpperModuleName}}Cache.FlushIndeof("{{TableName}}l")
			dao.{{UpperModuleName}}Cache.Flush("{{TableName}}c")
		`
	}
	rs += `
	c.JSON(200, ctx.R{Status: 1})
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
	// rs += generatedelete()
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
