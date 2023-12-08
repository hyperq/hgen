package template

const VueTemplete = `
<template>
  <div class="bg-white h-full">
    <BasicTable @register="registerTable">
      <template #toolbar>
        <Button type="primary" v-if="permission('maintain')" @click="edit({})">添加</Button>
      </template>
      <template #action="{ record }">
        <TableAction
          :actions="[
            {
              label: '编辑',
              color: 'success',
              onClick: edit.bind(null, record),
              ifShow: permission('maintain'),
            },
          ]"
        /> </template
    ></BasicTable>
    <CForm
      :info="editinfo"
      :visible="editflag"
      :saveapi="{{TableName}}.save"
      :schema="useModalSchemas()"
      title="维护{{TableComment}}"
      @cancel="editflag = false"
      @ok="reload()"
    />
  </div>
</template>

<script lang="ts" setup name="{{UpperTableName}}s">
  import { onActivated } from 'vue';
  import { Button } from 'ant-design-vue';
  import { TableAction, BasicTable } from '/@/components/Table';
  import { useColumns, useTableSchemas, useModalSchemas } from './model/_list';
  import useQuery from '/@/hooks/table/query';
  import permission from '/@/hooks/role/role';
  import { {{TableName}} } from '/@/api/{{TagName}}/{{TableName}}';
  import CForm from '@/components/common/Form.vue';

  const transPropsFunc = (e) => {
    if (!permission('maintain')) {
      e.actionColumn = undefined;
    }
  };

  const { registerTable, reload, edit, editinfo, editflag } = useQuery({
    api: {{TableName}},
    columns: useColumns(),
    schemas: useTableSchemas(),
    transPropsFunc,
    openEdit: true,
  });
  onActivated(async () => {
    reload();
  });
</script>

`

const VueModelString = `
import { BasicColumn } from '/@/components/Table/index';
import { ts2dt } from '/@/utils/time/time';
import { FormSchema } from '/@/components/Form/index';

export const useColumns = (): BasicColumn[] => {
  const columns: BasicColumn[] = [
    {{VueBasicColumns}}
  ];
  return columns;
};

export const useTableSchemas = (): FormSchema[] => {
  const schemas: FormSchema[] = [
    {
      field: 'custom.search',
      label: '搜索',
      component: 'InputTextArea',
      componentProps: {
        autoSize: true,
      },
    },
    {
      field: 'time.create_time',
      label: '创建时间',
      component: 'RangePicker',
      componentProps: {},
      colProps: { xs: 24, md: 24, lg: 16, xl: 12, xxl: 8 },
    },
  ];
  return schemas;
};

export const useModalSchemas = (): FormSchema[] => {
  const schemas: FormSchema[] = [
    {{VueForms}}
  ];
  return schemas;
};

`
