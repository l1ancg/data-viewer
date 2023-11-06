import * as React from 'react';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';

import { cn } from '@/lib/utils';
import ResourceEditor from '@/components/editor/resource-editor';

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Resource } from '@/types';
import { MixerHorizontalIcon } from '@radix-ui/react-icons';

let rs: Resource[] = [];
for (let i = 0; i < 10; i++) {
  rs.push({
    id: i,
    name: 'name-' + Math.floor(Math.random() * 1000000),
    type: 'mysql',
  });
}

let curr = 1;

export default function ResourcePanel() {
  function onSelected() {
    console.log('onSelected');
  }
  function onChange(e: React.MouseEvent<HTMLDivElement, MouseEvent>) {
    console.log('onChange');
    e.stopPropagation();
  }

  function editorDialog({ children }: { children: React.ReactNode }) {
    return (
      <Dialog>
        <DialogTrigger asChild>{children}</DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Data Resource Management</DialogTitle>
          </DialogHeader>
          <ResourceEditor></ResourceEditor>
          <DialogFooter>
            <Button type='submit'>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    );
  }

  return (
    <Card className='w-[250px] h-[400px]'>
      <CardHeader>
        <CardTitle>Resource</CardTitle>
      </CardHeader>
      <CardContent className='mb-0 pb-0'>
        <ScrollArea className='h-[280px]'>
          <div className='text-sm pr-3'>
            {rs.map((r) => (
              <div
                key={r.id}
                className={cn(
                  curr === r.id
                    ? 'border-primary bg-primary text-background'
                    : 'border-background border-dashed hover:border-dashed hover:border-primary',
                  'mt-2 mb-2 border pt-1 pb-1 pl-2 pr-2 rounded-lg hover:cursor-pointer flex flex-row'
                )}
                onClick={onSelected}
              >
                <div className='inline-block flex-1'>{r.name}</div>
                <div>
                  {editorDialog({
                    children: (
                      <MixerHorizontalIcon
                        className='w-5 h-5 inline-block hover:text-green-500'
                        onClick={() => onChange}
                      />
                    ),
                  })}
                </div>
              </div>
            ))}
          </div>
        </ScrollArea>
      </CardContent>
      <CardFooter className='pt-2'>
        {editorDialog({
          children: (
            <Button
              variant='outline'
              className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
            >
              Add resource
            </Button>
          ),
        })}
      </CardFooter>
    </Card>
  );
}
