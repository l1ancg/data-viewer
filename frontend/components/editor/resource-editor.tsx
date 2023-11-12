import { useState, FormEvent } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Resource } from '@/types';
import { baseMutate } from '@/lib/graphql';

export interface IResourceEditorProps {
  resource: Resource | null;
  onRefresh: () => void;
}

export default function ResourceEditor({
  resource,
  onRefresh,
}: IResourceEditorProps) {
  if (!resource) {
    resource = {
      name: '',
      type: '',
      data: '',
    };
  }
  const [data, setData] = useState(resource);
  const handleChange = (name: string, value: string) => {
    setData({
      ...data,
      [name]: value,
    });
  };

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    if (!data.name || !data.type || !data.data) {
      alert('所有字段都是必填的！');
      return;
    }

    const mutation = `
        mutation Save($id: Int, $name: String!, $type: String!, $data: String!) {
          resource(id: $id, name: $name, type: $type, data: $data) {
            id
            name
            type
            data
          }
        }
    `;
    const variables = {
      id: resource?.id,
      name: data.name,
      type: data.type,
      data: data.data,
    };
    baseMutate(mutation, variables)
      .then(() => onRefresh())
      .catch((e) => console.log(e));
  };

  return (
    <>
      <Dialog defaultOpen onOpenChange={() => onRefresh()}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Data Resource Management</DialogTitle>
          </DialogHeader>
          <form className='grid gap-4 py-4' onSubmit={handleSubmit}>
            <div className='grid grid-cols-5 items-center gap-4'>
              <Label htmlFor='name' className='text-right'>
                Name
              </Label>
              <Input
                className='col-span-4'
                value={data.name}
                onChange={(e) => handleChange('name', e.target.value)}
                required
              />
            </div>
            <div className='grid grid-cols-5 items-center gap-4'>
              <Label htmlFor='name' className='text-right'>
                Type
              </Label>
              <Select
                value={data.type}
                onValueChange={(v) => handleChange('type', v)}
                required
              >
                <SelectTrigger className='col-span-4'>
                  <SelectValue placeholder='Select a resource type' />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <SelectItem value='mysql'>MySQL</SelectItem>
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
            <div className='grid grid-cols-5 items-center gap-4'>
              <Label htmlFor='username' className='text-right'>
                Data
              </Label>
              <Textarea
                className='col-span-4'
                value={data.data}
                onChange={(e) => handleChange('data', e.target.value)}
              />
            </div>
            <Button type='submit'>Save</Button>
          </form>
        </DialogContent>
      </Dialog>
    </>
  );
}
