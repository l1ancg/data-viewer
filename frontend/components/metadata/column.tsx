import { toast } from '@/components/ui/use-toast';
import { MyField } from '@/components/editor/editor';
import { Column } from '@/types';

const Fields: Array<MyField> = [
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
  {
    name: 'desc',
    type: 'input',
  },
];

const Query = `
query GetColumns($resourceId: Int!) {
  columns(resourceId: $resourceId) {
    id
    resourceId
    dictId
    name
    dataType
    orderBy
    display
    condition
    desc
  }
}
`;

const Delete = ``;

// graphql mutation
const Save = `
mutation Save(
  $id: Int
  $resourceId: Int!
  $dictId: Int
  $name: String!
  $dataType: String!
  $orderBy: String!
  $display: Boolean!
  $condition: Boolean!
  $desc: String!
) {
  createColumn(
    id: $id
    resourceId: $resourceId
    dictId: $dictId
    name: $name
    dataType: $dataType
    orderBy: $orderBy
    display: $display
    condition: $condition
    desc: $desc
  ) {
    id
    resourceId
    dictId
    name
    dataType
    orderBy
    display
    condition
    desc
  }
}
`;

const OnValidate = (row: Column) => {
  if (!row.resourceId) {
    toast({ title: 'Please select a resource' });
  }
  if (!row.name) {
    toast({ title: 'Please input a name' });
  }
  if (!row.dataType) {
    toast({ title: 'Please select a data type' });
  }
};

const NewData = (resourceId: number): Column => {
  return {
    resourceId: resourceId,
    name: '',
    dataType: '',
    orderBy: '',
    display: false,
    condition: false,
    desc: '',
  };
};

export default { Fields, Query, Save, Delete, OnValidate, NewData };
