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
import { MouseEvent, useEffect, useState } from 'react';
import { useToast } from '@/components/ui/use-toast';
import { MyEditor } from '@/components/editor/editor';
import columnMetadata from '@/components/metadata/column';
import { baseQuery } from '@/lib/graphql';

export interface ColumnPanelProps {
  resourceId: number | null;
}

export default function ColumnPanel({ resourceId }: ColumnPanelProps) {
  const [columns, setColumns] = useState<Column[]>([]);
  const [column, setColumn] = useState<Column | null>(null);
  const [openEditor, setOpenEditor] = useState<boolean>(false);
  const { toast } = useToast();

  const deleteHandle = (event: MouseEvent<SVGElement>, column: Column) => {
    event.preventDefault();
  };
  const editHandle = (event: MouseEvent<SVGElement>, column: Column) => {
    event.preventDefault();
    setColumn(column);
    setOpenEditor(true);
  };

  const addNewHandle = () => {
    if (!resourceId) {
      toast({
        variant: 'destructive',
        title: 'Please select a resource first',
      });
      return;
    }
    setColumn(columnMetadata.NewData(resourceId));
    setOpenEditor(true);
  };
  const closeEditorHandle = () => {
    setOpenEditor(false);
    setColumn(null);
    fetchColumns();
  };

  const fetchColumns = () => {
    if (resourceId) {
      baseQuery<{ columns: Column[] }>(columnMetadata.Query, { resourceId })
        .then((data) => setColumns(data.columns))
        .catch((e) => toast({ variant: 'destructive', title: e.message }));
    } else {
      setColumns([]);
    }
  };

  useEffect(() => {
    fetchColumns();
  }, [resourceId]);

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
                            variant={column.display ? 'default' : 'outline'}
                            className='mr-2'
                          >
                            Display
                          </Badge>
                          <Badge
                            variant={column.orderBy ? 'default' : 'outline'}
                            className='mr-2'
                          >
                            Order
                          </Badge>
                          <Badge
                            variant={column.condition ? 'default' : 'outline'}
                            className='mr-2'
                          >
                            Condition
                          </Badge>
                        </div>
                      </TableCell>
                      <TableCell>
                        <div className='flex flex-row'>
                          <Pencil1Icon
                            className='h-4 w-4 mr-2 hover:text-green-500 hover:cursor-pointer'
                            onClick={(e) => editHandle(e, column)}
                          />
                          <MinusCircledIcon
                            className='h-4 w-4 hover:text-red-500 hover:cursor-pointer'
                            onClick={(e) => deleteHandle(e, column)}
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
          <Button
            variant='outline'
            className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
            onClick={addNewHandle}
          >
            Add column
          </Button>
        </CardFooter>
      </Card>
      {openEditor && (
        <MyEditor
          row={column}
          mutate={columnMetadata.Save}
          fields={columnMetadata.Fields}
          onRefresh={closeEditorHandle}
          onValidate={columnMetadata.OnValidate}
        />
      )}
    </>
  );
}
