import * as React from 'react';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectSeparator,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { DataTypeValues } from '@/types';

import { Toggle } from '@/components/ui/toggle';
export interface IColumnEditorProps {}

export default function ColumnEditor(props: IColumnEditorProps) {
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
            Label
          </Label>
          <Input id='name1' className='col-span-3' />
        </div>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='name' className='text-right'>
            Type
          </Label>
          <Select>
            <SelectTrigger className='col-span-3'>
              <SelectValue placeholder='Select a data type' />
            </SelectTrigger>
            <SelectContent>
              {DataTypeValues.map((dt) => (
                <SelectItem
                  key={dt.value}
                  value={dt.value}
                  className='cursor-pointer'
                >
                  {dt.label}
                </SelectItem>
              ))}
              <SelectSeparator></SelectSeparator>
              <SelectItem
                value='add'
                className='text-slate-400 border-slate-400 border-2 border-dashed cursor-pointer'
              >
                Add
              </SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div className='grid grid-cols-4 items-center gap-4'>
          <Label htmlFor='username' className='text-right'>
            Data
          </Label>
          <div className='col-span-3'>
            <Toggle
              aria-label='Toggle italic'
              onPressedChange={(x) => console.log(x)}
            >
              <div>Display</div>
            </Toggle>
            <Toggle
              aria-label='Toggle italic'
              onPressedChange={(x) => console.log(x)}
            >
              <div>Asc</div>
            </Toggle>
            <Toggle
              aria-label='Toggle italic'
              onPressedChange={(x) => console.log(x)}
            >
              <div>Desc</div>
            </Toggle>
            <Toggle
              aria-label='Toggle italic'
              onPressedChange={(x) => console.log(x)}
            >
              <div>Condition</div>
            </Toggle>
          </div>
        </div>
      </div>
    </>
  );
}
