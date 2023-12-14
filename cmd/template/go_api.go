package template

const ApiString = `
package {{TagName}}

// func {{UpperTableName}}(c *ctx.Context) {
// 	q := qs.New("a.id=?", c.Param("id"))
// 	q.JoinSelect = []string{}
// 	q.Join = []string{}
// 	data, err := dao.Find[model.{{UpperTableName}}](q)
// 	if c.HandlerError(err) {
// 		return
// 	}
// 	c.RespJSON(data)
// }

// func {{UpperTableName}}s(c *ctx.Context) {
// 	q := qs.Auto(c)
// 	q.JoinSelect = []string{}
// 	q.Join = []string{}
// 	q.CustomSearchByQP(c.Query("qp"))
// 	pd, err := dao.FindsByPage[model.{{UpperTableName}}](q)
// 	if c.HandlerError(err) {
// 		return
// 	}
// 	c.RespJSON(pd)
// }


// func {{UpperTableName}}Save(c *ctx.Context) {
// 	pd, err := ctx.UnmarshalFromStringT[model.{{UpperTableName}}](c)
// 	if c.HandlerError(err) {
// 		return
// 	}
// 	userid := c.GetID()
// 	var op model.OperateEnum
// 	opd := model.{{UpperTableName}}{}
// 	if pd.ID > 0 {
// 		op = model.Operate.Update
// 		opd, err = dao.FindByIDCache[model.{{UpperTableName}}](pd.ID)
// 		if c.HandlerError(err) {
// 			return
// 		}
// 		if !c.VersionCheck(pd.Version, opd.Version) {
// 			return
// 		}
// 		pd.ModifyBy = userid
// 	} else {
// 		op = model.Operate.Create
// 		pd.CreateBy = userid
// 	}
// 	pd.Version++
// 	oid, err := dao.InsertOrUpdate(&pd)
// 	if c.HandlerError(err) {
// 		return
// 	}
// 	id := oid.(int64)
// 	dao.ClearCacheT[model.{{UpperTableName}}](id)
// 	operates.CreateOperateRecord[model.{{UpperTableName}}](op, pd, opd, c.GetID(), id, "")
// 	c.RespJSON(true)
// }
	

// g.GET("/{{TableName}}/:id", ctx.H({{UpperTableName}}))
// g.GET("/{{TableName}}s", ctx.H({{UpperTableName}}s))
// g.POST("/{{TableName}}", ctx.H({{UpperTableName}}Save))

// g.GET("/{{TableName}}/:id", ctx.H(cs.One[model.{{UpperTableName}}]))
// g.GET("/{{TableName}}s", ctx.H(cs.List[model.{{UpperTableName}}]))
// g.POST("/{{TableName}}", ctx.H(cs.Save[model.{{UpperTableName}}]))
// g.DELETE("/{{TableName}}", ctx.H(cs.Del[model.{{UpperTableName}}]))
`
