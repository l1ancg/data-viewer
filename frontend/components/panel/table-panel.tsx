import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { useEffect, useState } from 'react';
import { Column } from '@/types';

export interface TablePanelProps {
  data: Array<Map<string, object>>;
  columns: Column[];
}

export default function TablePanel(props: TablePanelProps) {
  const { data, columns } = props;

  return (
    <>
      <Table>
        <TableHeader>
          <TableRow>
            {
            columns.map((column) => (
              <TableCell key={column.name}>{column.label}</TableCell>
            ))
            }
            
          </TableRow>
        </TableHeader>
        <TableBody>
          {data.map((item) => (
            // @ts-ignore
            <TableRow key={item.name}>
              {columns.map((column) => (
                // @ts-ignore
                <TableCell key={column.name}>{item[column.name]}</TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </>
  );
}
