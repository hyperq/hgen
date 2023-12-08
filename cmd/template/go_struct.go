package template

const ModelString = `
package model

// {{UpperTableName}} {{TableComment}}
type {{UpperTableName}} struct{
	BaseStruct
	{{TableColumns}}
}

func (b {{UpperTableName}}) TableName() string {
	return "{{TableName}}"
}
`
