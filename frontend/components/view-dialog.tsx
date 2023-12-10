import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import ResourcePanel from '@/components/panel/resource-panel';
import OptionPanel from '@/components/panel/option-panel';
import SqlEditor from '@/components/editor/sql-editor';
import { useState, ReactNode } from 'react';
import parse from '@/lib/sql';
import { Option, Paramter, Column } from '@/types';

export function ViewDialog({ children }: { children: ReactNode }) {
  const [selectedResourceId, setSelectedResourceId] = useState<number | null>(
    null
  );
  const [sql, setSql] = useState('');
  const [option, setOption] = useState<Option>({
    columns: [],
    parameters: [],
    version: 0,
  });

  const inputedSqlHandle = (s: string) => {
    if (s === sql) {
      return;
    }
    if (!s) {
      setSql('');
      setOption({ columns: [], parameters: [], version: 0 });
      return;
    }
    const pr = parse(s, 'MySQL');
    setSql(pr.sql || '');

    let parameters: Paramter[] = [];
    let columns: Column[] = [];
    if (pr.whereList) {
      pr.whereList.forEach((w) => {
        parameters.push({ label: w, name: w });
      });
    }
    if (pr.columnList) {
      pr.columnList.forEach((c) => {
        columns.push({ label: c, name: c });
      });
    }
    setOption({ columns, parameters, version: 0 });
  };
  const editedOptionHandle = (newOption: Option) => {
    console.log(newOption);
    if (option.version != newOption.version) {
      setOption(newOption);
    }
  };
  const saveHandle = () => {
    console.log(sql);
    console.log(option);
  };

  return (
    <Dialog onOpenChange={() => setSelectedResourceId(null)}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-w-max'>
        <DialogHeader>
          <DialogTitle>View Editor</DialogTitle>
        </DialogHeader>
        <div className='gap-1 flex flex-row'>
          <div className=''>
            <ResourcePanel
              onSelectedResource={setSelectedResourceId}
            ></ResourcePanel>
          </div>
          <div className=''>
            <SqlEditor onInputedSql={inputedSqlHandle}></SqlEditor>
          </div>
          <div className=''>
            <OptionPanel
              option={option}
              onEditedOption={editedOptionHandle}
            ></OptionPanel>
          </div>
        </div>
        <DialogFooter>
          <Button onClick={saveHandle}>Save</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
