import * as React from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';

export interface IMysqlEditorProps {}

export default function MysqlEditor(props: IMysqlEditorProps) {
  return (
    <>
      <div className='grid gap-4 py-4'>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='name' className='text-right'>
            Name
          </Label>
          <Input id='name' className='col-span-3' />
        </div>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='name' className='text-right'>
            Type
          </Label>
          <Input id='name' className='col-span-3' />
        </div>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='username' className='text-right'>
            Data
          </Label>
          <Input id='username' className='col-span-3' />
        </div>
      </div>
    </>
  );
}
