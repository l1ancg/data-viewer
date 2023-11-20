import { toast } from '@/components/ui/use-toast';
import { MyField } from '@/components/editor/editor';
import { Resource } from '@/types';

const Fields: Array<MyField> = [
  {
    name: 'name',
    type: 'input',
  },
  {
    name: 'type',
    type: 'select',
    options: [
      { value: 'mysql', label: 'mysql' },
      { value: 'mongodb', label: 'mongodb' },
    ],
  },
  {
    name: 'data',
    type: 'textarea',
  },
];

const Query = `
{
  resources {
    id
    name
    type
    data
  }
}
`;

const Delete = ``;

// graphql mutation
const Save = `
mutation Save($id: Int, $name: String!, $type: String!, $data: String!) {
  resource(id: $id, name: $name, type: $type, data: $data) {
    id
    name
    type
    data
  }
}`;

const OnValidate = (row: Resource) => {
  if (!row.name) {
    toast({ variant: 'destructive', title: 'Please input a name' });
  }
  if (!row.type) {
    toast({ variant: 'destructive', title: 'Please select a type' });
  }
  if (!row.data) {
    toast({ variant: 'destructive', title: 'Please input a data' });
  }
};

const NewData = (): Resource => {
  return {
    name: '',
    type: '',
    data: '',
  };
};

export default { Fields, Query, Save, Delete, OnValidate, NewData };
