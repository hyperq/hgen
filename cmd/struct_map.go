package cmd

type columnType struct {
	TransferType   string
	TransferInsert func(string) string
}

var sqltogotype = map[string]columnType{
	"tinyint": columnType{
		TransferType: "int",
	},
	"smallint": columnType{
		TransferType: "int",
	},
	"mediumint": columnType{
		TransferType: "int",
	},
	"int": columnType{
		TransferType: "int",
	},
	"integer": columnType{
		TransferType: "int",
	},
	"bigint": columnType{
		TransferType: "int64",
	},
	"float": columnType{
		TransferType: "float64",
	},
	"double": columnType{
		TransferType: "float64",
	},
	"decimal": columnType{
		TransferType: "float64",
	},
	"date": columnType{
		TransferType: "time.Time",
	},
	"time": columnType{
		TransferType: "string",
	},
	"year": columnType{
		TransferType: "int",
	},
	"datetime": columnType{
		TransferType: "time.Time",
	},
	"timestamp": columnType{
		TransferType: "int",
	},
	"datetimeoffset": columnType{
		TransferType: "datetime",
	},
	"char": columnType{
		TransferType: "string",
	},
	"varchar": columnType{
		TransferType: "string",
	},
	"tinyblob": columnType{
		TransferType: "string",
	},
	"tinytext": columnType{
		TransferType: "string",
	},
	"blob": columnType{
		TransferType: "string",
	},
	"text": columnType{
		TransferType: "string",
	},
	"mediumblob": columnType{
		TransferType: "string",
	},
	"mediumtext": columnType{
		TransferType: "string",
	},
	"longblob": columnType{
		TransferType: "string",
	},
	"longtext": columnType{
		TransferType: "string",
	},
}

var sqltotstype = map[string]columnType{
	"tinyint": columnType{
		TransferType: "number",
	},
	"smallint": columnType{
		TransferType: "number",
	},
	"mediumint": columnType{
		TransferType: "number",
	},
	"int": columnType{
		TransferType: "number",
	},
	"integer": columnType{
		TransferType: "number",
	},
	"bigint": columnType{
		TransferType: "number",
	},
	"float": columnType{
		TransferType: "number",
	},
	"double": columnType{
		TransferType: "number",
	},
	"decimal": columnType{
		TransferType: "number",
	},
	"date": columnType{
		TransferType: "string",
	},
	"time": columnType{
		TransferType: "string",
	},
	"year": columnType{
		TransferType: "number",
	},
	"datetime": columnType{
		TransferType: "string",
	},
	"timestamp": columnType{
		TransferType: "number",
	},
	"datetimeoffset": columnType{
		TransferType: "string",
	},
	"char": columnType{
		TransferType: "string",
	},
	"varchar": columnType{
		TransferType: "string",
	},
	"tinyblob": columnType{
		TransferType: "string",
	},
	"tinytext": columnType{
		TransferType: "string",
	},
	"blob": columnType{
		TransferType: "string",
	},
	"text": columnType{
		TransferType: "string",
	},
	"mediumblob": columnType{
		TransferType: "string",
	},
	"mediumtext": columnType{
		TransferType: "string",
	},
	"longblob": columnType{
		TransferType: "string",
	},
	"longtext": columnType{
		TransferType: "string",
	},
}
