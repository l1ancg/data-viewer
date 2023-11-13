import { Column, DataTypeValues } from '@/types';

import { MyEditor, MyField } from './editor';
import { useState } from 'react';

export interface IColumnEditorProps {
  data: Column | null;
  resourceId: number;
  onRefresh: () => void;
}
// todo 提出去 && 删除此文件放到panel里面
const fields: Array<MyField> = [
  {
    name: 'name',
    type: 'input',
  },
  {
    name: 'dataType',
    type: 'select',
    options: [
      { value: 'string', label: 'string' },
      { value: 'number', label: 'number' },
    ],
  },
  {
    name: 'orderBy',
    type: 'select',
    options: [
      { value: 'asc', label: 'asc' },
      { value: 'desc', label: 'desc' },
    ],
  },
  {
    name: 'display',
    type: 'switch',
  },
  {
    name: 'condition',
    type: 'switch',
  },
];

const mutate = `
mutation Save($id: Int, $viewId: Int!, $dictId: Int!, $name: String!, $dataType: String!, $orderBy: String!, $display: Boolean!, $condition: Boolean!) {
  view(
    id: $id
    viewId: $viewId
    dictId: $dictId
    name: $name
    dataType: $dataType
    orderBy: $orderBy
    display: $display
    condition: $condition
  ) {
    id
    resourceId
    resourceType
    displayType
    name
    desc
  }
}`;

export default function ColumnEditor(props: IColumnEditorProps) {
  const { data, resourceId, onRefresh } = props;
  const [column, setColumn] = useState<Column | null>(data);
  if (!column) {
    setColumn({
      resourceId: resourceId,
      name: '',
      dataType: '',
      orderBy: '',
      display: false,
      condition: false,
    });
  }

  return (
    <>
      <MyEditor
        row={column}
        mutate={mutate}
        fields={fields}
        onRefresh={onRefresh}
      />
    </>
  );
}
