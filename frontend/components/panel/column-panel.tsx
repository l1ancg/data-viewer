import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { ScrollArea } from '@/components/ui/scroll-area';
import { MinusCircledIcon, Pencil1Icon } from '@radix-ui/react-icons';
import { Column } from '@/types';
import ColumnEditor from '@/components/editor/column-editor';
import { MouseEvent, useEffect, useState } from 'react';

export default function ColumnPanel() {
  const [columns, setColumns] = useState<Column[]>([]);
  const [column, setColumn] = useState<Column | null>(null);
  const [openEditor, setOpenEditor] = useState<boolean>(false);

  const handleDelete = (
    event: MouseEvent<HTMLButtonElement>,
    column: Column
  ) => {
    event.preventDefault();
  };
  const handleEdit = (event: MouseEvent<HTMLButtonElement>, column: Column) => {
    event.preventDefault();
    setColumn(column);
    setOpenEditor(true);
  };

  const fetchColumns = () => {};

  useEffect(() => {});

  return (
    <>
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
                      <TableCell className='font-medium'>
                        {column.name}
                      </TableCell>
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
                          <Button
                            asChild
                            onClick={(e) => handleDelete(e, column)}
                          >
                            <Pencil1Icon className='h-4 w-4 mr-2 hover:text-green-500 hover:cursor-pointer' />
                          </Button>
                          <Button
                            asChild
                            onClick={(e) => handleEdit(e, column)}
                          >
                            <MinusCircledIcon className='h-4 w-4 hover:text-red-500 hover:cursor-pointer' />
                          </Button>
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
          <Button
            variant='outline'
            className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
            onClick={() => setOpenEditor(true)}
          >
            Add column
          </Button>
        </CardFooter>
      </Card>
      {openEditor && (
        <ColumnEditor data={column} resourceId={0} onRefresh={fetchColumns} />
      )}
    </>
  );
}
