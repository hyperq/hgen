package template

const TsApiString = `
import { BaseStruct, BaseApi } from '../base/api';

export interface {{UpperTableName}} extends BaseStruct {
  {{TsInterfaces}}
}

export const {{TableName}} = new BaseApi<{{UpperTableName}}>('{{TableName}}');
`
