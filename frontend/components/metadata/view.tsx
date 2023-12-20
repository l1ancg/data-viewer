import { toast } from '@/components/ui/use-toast';
import { View } from '@/types';

const Query = `
query views {
  views {
    id
    name
    resourceId
    ql
    options
  }
}
`;

const Delete = `
  mutation view($id: Int!) {
    delete(id: $id) {
      id
    }
  }
`;

// graphql mutation
const Save = `
  mutation view(
    $id: Int
    $name: String
    $resourceId: Int
    $ql: String
    $options: String
    ) {
    view(
      id: $id
      name: $name
      resourceId: $resourceId
      ql: $ql
      options: $options
    ) {
      id
      name
      resourceId
      ql
      options
    }
  }
`;

const OnValidate = (row: View) => {
  if (!row.resourceId) {
    toast({ variant: 'destructive', title: 'Please select a Resource' });
    return false;
  }
  if (!row.name) {
    toast({ variant: 'destructive', title: 'Please input a name' });
    return false;
  }
  if (!row.ql) {
    toast({ variant: 'destructive', title: 'Please input a ql' });
    return false;
  }
  return true;
};

export default { Query, Save, Delete, OnValidate };
