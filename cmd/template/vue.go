package template

const VueApiString = `
import { defHttp } from '/@/utils/http/axios'
import type { FetchParam, RespList } from '/#/axios'

import { {{UpperTableName}} } from './model/{{TableName}}'

export function {{TableName}}s(params: FetchParam) {
  return defHttp.get<RespList<{{UpperTableName}}>>({
    url: '{{TableName}}s',
    params,
  })
}

export function {{TableName}}(id: string | number) {
  return defHttp.get<{{UpperTableName}}>({
    url: '{{TableName}}/' + id,
  })
}

export function {{TableName}}save(params: {{UpperTableName}}) {
  return defHttp.post<{{UpperTableName}}>({
    url: '{{TableName}}',
	params,
  })
}`

const VueModelString = `
export interface {{UpperTableName}} {
	{{VueModel}}
}
`
