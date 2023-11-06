import * as React from 'react';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Textarea } from '@/components/ui/textarea';

export interface IResourceEditorProps {}

export default function ResourceEditor(props: IResourceEditorProps) {
  return (
    <>
      <div className='grid gap-4 py-4'>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='name' className='text-right'>
            Name
          </Label>
          <Input id='name' className='col-span-3' />
        </div>{' '}
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='name' className='text-right'>
            Type
          </Label>
          <Select>
            <SelectTrigger className='col-span-3'>
              <SelectValue placeholder='Select a resource type' />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value='mysql'>MySQL</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
        </div>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='username' className='text-right'>
            Data
          </Label>
          <Textarea
            className='col-span-3'
            placeholder='Type your message here.'
          />
        </div>
      </div>
    </>
  );
}
