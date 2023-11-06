import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover';

import * as React from 'react';

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Badge } from '@/components/ui/badge';
import { ScrollArea } from '@/components/ui/scroll-area';
import { MinusCircledIcon, Pencil1Icon } from '@radix-ui/react-icons';
import { Column } from '@/types';
import ColumnEditor from '@/components/editor/column-editor';
import { on } from 'events';

const columns: Column[] = [];
for (let i = 0; i < 10; i++) {
  columns.push({
    id: i,
    name: 'name' + i,
    label: 'name' + i,
    dataType: 'mysql',
    display: Math.random() < 0.5,
    orderBy: Math.random() < 0.5,
    condition: Math.random() < 0.5,
  });
}

export default function ColumnPanel() {
  function editorDialog({ children }: { children: React.ReactNode }) {
    return (
      <Dialog>
        <DialogTrigger asChild>{children}</DialogTrigger>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Data Resource Management</DialogTitle>
          </DialogHeader>
          <ColumnEditor></ColumnEditor>
          <DialogFooter>
            <Button type='submit'>Save</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    );
  }

  function onEdit(col: Column) {
    console.log('edit column:', col);
  }
  function onRemove(col: Column) {
    console.log('remove column:', col);
  }

  return (
    <Card className='w-[600px] h-[400px]'>
      <CardHeader>
        <CardTitle>Column</CardTitle>
      </CardHeader>
      <CardContent className='pb-0 mb-0'>
        <ScrollArea className='h-[280px]'>
          <div className='pr-3'>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Name</TableHead>
                  <TableHead>Type</TableHead>
                  <TableHead>Property</TableHead>
                  <TableHead className='w-[30px] pr-3'>Option</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {columns.map((column) => (
                  <TableRow key={column.name}>
                    <TableCell className='font-medium'>{column.name}</TableCell>
                    <TableCell>{column.dataType}</TableCell>
                    <TableCell>
                      <div className=''>
                        <Badge
                          variant={column.display ? 'outline' : 'default'}
                          className='mr-2'
                        >
                          Display
                        </Badge>
                        <Badge
                          variant={column.orderBy ? 'outline' : 'default'}
                          className='mr-2'
                        >
                          Order
                        </Badge>
                        <Badge
                          variant={column.condition ? 'outline' : 'default'}
                          className='mr-2'
                        >
                          Condition
                        </Badge>
                      </div>
                    </TableCell>
                    <TableCell>
                      <div className='flex flex-row'>
                        {editorDialog({
                          children: (
                            <Pencil1Icon
                              className='h-4 w-4 mr-2 hover:text-green-500 hover:cursor-pointer'
                              onClick={() => onEdit(column)}
                            />
                          ),
                        })}
                        <MinusCircledIcon
                          className='h-4 w-4 hover:text-red-500 hover:cursor-pointer'
                          onClick={() => onRemove(column)}
                        />
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
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
              Add column
            </Button>
          ),
        })}
      </CardFooter>
    </Card>
  );
}
