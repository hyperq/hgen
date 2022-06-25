package template

const ApiString = `
package {{TableName}}

import (
	"backend/model"
)

// {{OneTableName}}
// @tags		{{TagName}}
// @Summary 	根据id获取{{TableName}}信息
// @Description 根据id获取{{TableName}}信息
// @Accept  	json
// @Produce 	json
// @Param 		id query string true "{{TableName}} id"
// @Success 	200 {object} model.{{UpperTableName}}
// @Failure 	400 {object} ctx.R
// @Router 		/api/v1/{{TableName}} [get]
func {{OneTableName}}(c *ctx.Context) {
	var data model.{{UpperTableName}}
	err := dao.FindByIDCache(c.Query("id"),&data)
	if c.HandlerError(err) {
		return
	}
	c.HandlerOk(data)
}


// {{ListTableName}}
// @tags 		{{TagName}}
// @Summary 	获取{{TableName}}列表
// @Description 获取{{TableName}}列表
// @Accept  	json
// @Produce 	json
// @Success 	200 {array} model.{{UpperTableName}}
// @Failure 	400 {object} ctx.R
// @Router 		/api/v1/{{TableName}}s [get]
func {{UpperTableName}}s (c *ctx.Context) {
	q := qs.Auto(c)
	var data []model.{{UpperTableName}}
	err := dao.FindsCache(q,&data)
	if c.HandlerError(err) {
		return
	}
	c.HandlerOk(gin.H{"data": data, "total": dao.CountCache("{{TableName}}", q)})
}

// {{SaveTableName}}
// @tags 		{{TagName}}
// @Summary 	添加或更新{{TableName}}
// @Description 添加或更新{{TableName}}
// @Accept  	x-www-form-urlencoded
// @Produce 	json
// @Param 		document body model.{{UpperTableName}} true "{{TableName}}信息"
// @Success 	200 {number} int
// @Failure 	400 {object} ctx.R
// @Router 		/api/v1/{{TableName}} [post]
func {{SaveTableName}}(c *ctx.Context) {
	// 获取用户id
	var pd model.{{UpperTableName}}
	err := c.UnmarshalFromString(&pd)
	if c.HandlerError(err) {
		return
	}
	// 获取原来的数据
	var opd model.{{UpperTableName}}
	if pd.ID != 0 {
		err = dao.FindByIDCache(pd.ID,&opd)
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
	// 清空相关缓存
	dao.ClearCache("{{TableName}}",id)
	c.HandlerOk(id)
}

// {{TableName}}
//g.GET("/{{TableName}}", ctx.H(rs.{{OneTableName}}))
//g.GET("/{{TableName}}s", ctx.H(rs.{{ListTableName}}))
//g.POST("/{{TableName}}", ctx.H(rs.{{SaveTableName}}))
//g.GET("/{{TableName}}", cs.M(reflect.TypeOf(model.{{UpperTableName}}{})), ctx.H(cs.Get))
//g.GET("/{{TableName}}s",cs.M(reflect.TypeOf([]{{model.{{UpperTableName}}{})), ctx.H(cs.Gets))
//g.POST("/{{TableName}}",cs.M(reflect.TypeOf({{model.{{UpperTableName}}{})), ctx.H(cs.Save))
`
