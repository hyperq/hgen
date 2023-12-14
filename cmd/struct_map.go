package cmd

type columnType struct {
	TransferType   string
	TransferInsert func(string) string
}

var Sql2GoType = map[string]columnType{
	"tinyint": {
		TransferType: "bool",
	},
	"smallint": {
		TransferType: "int64",
	},
	"mediumint": {
		TransferType: "int64",
	},
	"int": {
		TransferType: "int64",
	},
	"integer": {
		TransferType: "int64",
	},
	"bigint": {
		TransferType: "int64",
	},
	"float": {
		TransferType: "float64",
	},
	"double": {
		TransferType: "float64",
	},
	"decimal": {
		TransferType: "float64",
	},
	"date": {
		TransferType: "time.Time",
	},
	"time": {
		TransferType: "string",
	},
	"year": {
		TransferType: "int",
	},
	"datetime": {
		TransferType: "time.Time",
	},
	"timestamp": {
		TransferType: "int",
	},
	"datetimeoffset": {
		TransferType: "datetime",
	},
	"char": {
		TransferType: "string",
	},
	"varchar": {
		TransferType: "string",
	},
	"tinyblob": {
		TransferType: "string",
	},
	"tinytext": {
		TransferType: "string",
	},
	"blob": {
		TransferType: "string",
	},
	"text": {
		TransferType: "string",
	},
	"mediumblob": {
		TransferType: "string",
	},
	"mediumtext": {
		TransferType: "string",
	},
	"longblob": {
		TransferType: "string",
	},
	"longtext": {
		TransferType: "string",
	},
}

var Sql2TsType = map[string]columnType{
	"tinyint": {
		TransferType: "boolean",
	},
	"smallint": {
		TransferType: "number",
	},
	"mediumint": {
		TransferType: "number",
	},
	"int": {
		TransferType: "number",
	},
	"integer": {
		TransferType: "number",
	},
	"bigint": {
		TransferType: "number",
	},
	"float": {
		TransferType: "number",
	},
	"double": {
		TransferType: "number",
	},
	"decimal": {
		TransferType: "number",
	},
	"date": {
		TransferType: "string",
	},
	"time": {
		TransferType: "string",
	},
	"year": {
		TransferType: "number",
	},
	"datetime": {
		TransferType: "string",
	},
	"timestamp": {
		TransferType: "number",
	},
	"datetimeoffset": {
		TransferType: "string",
	},
	"char": {
		TransferType: "string",
	},
	"varchar": {
		TransferType: "string",
	},
	"tinyblob": {
		TransferType: "string",
	},
	"tinytext": {
		TransferType: "string",
	},
	"blob": {
		TransferType: "string",
	},
	"text": {
		TransferType: "string",
	},
	"mediumblob": {
		TransferType: "string",
	},
	"mediumtext": {
		TransferType: "string",
	},
	"longblob": {
		TransferType: "string",
	},
	"longtext": {
		TransferType: "string",
	},
}

var Sql2VueFormType = map[string]columnType{
	"tinyint": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'RadioButtonGroup',
			componentProps: {
				options: [
					{ label: '是', value: true },
					{ label: '否', value: false },
				],
			},
		}
		`,
	},
	"smallint": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 0,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"mediumint": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 0,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"int": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 0,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"integer": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 0,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"bigint": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 0,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"float": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 2,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"double": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 2,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"decimal": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 2,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"date": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'Input',
			rules: [{ required: false }],
		}`,
	},
	"time": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'Input',
			rules: [{ required: false }],
		}`,
	},
	"year": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 2,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"datetime": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'Input',
			rules: [{ required: false }],
		}`,
	},
	"timestamp": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputNumber',
			componentProps: () => {
				return {
					min: 0,
					precision: 2,
				};
			},
			rules: [{ required: false }],
		}`,
	},
	"datetimeoffset": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'Input',
			rules: [{ required: false }],
		}`,
	},
	"char": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'Input',
			rules: [{ required: false }],
		}`,
	},
	"varchar": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"tinyblob": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"tinytext": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"blob": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"text": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"mediumblob": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"mediumtext": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"longblob": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
	"longtext": {
		TransferType: `
		{
			field: '{{field}}',
			label: '{{label}}',
			component: 'InputTextArea',
			componentProps: {
				autoSize: true,
			},
			rules: [{ required: false }],
		}`,
	},
}

const VueFormTransID = `
{
	field: '{{field}}',
	label: '{{label}}',
	component: 'ApiSelect',
    componentProps: () => {
        return {
			api: .list,
			params: { current: -1 },
			resultField: 'data',
			labelField: 'name',
			valueField: 'id',
        };
    },
    rules: [{ required: true, type: 'number' }],
}`
const VueFormTransTime = `
{
	field: '{{field}}',
	label: '{{label}}',
	component: 'DatePickerTs',
	componentProps: () => {
		return {
			format: 'YYYY-MM-DD HH:mm:ss',
			showTime: {},
        };
    },
    rules: [{ required: true, type: 'number' }],
}`
