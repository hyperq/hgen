package template

const ModelString = `
package model

import (
	"time"
)

// {{UpperTableName}} {{TableComment}}
type {{UpperTableName}} struct{
	{{TableColumns}}
}

func (b {{UpperTableName}}) TableName() string {
	return "{{TableName}}"
}
`
