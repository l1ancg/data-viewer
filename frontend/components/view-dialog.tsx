import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { useToast } from '@/components/ui/use-toast';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import ResourcePanel from '@/components/panel/resource-panel';
import OptionPanel from '@/components/panel/option-panel';
import SqlEditor from '@/components/editor/sql-editor';
import { useState, ReactNode, useEffect } from 'react';
import parse from '@/lib/sql';
import { Option, Paramter, Column, View } from '@/types';
import viewMetadata from '@/components/metadata/view';
import { baseMutate } from '@/lib/graphql';
import { setReactionScheduler } from 'mobx/dist/internal';

export interface ViewDialogProps {
  children: ReactNode;
  view: View | null;
  open: boolean;
  onClose: (refresh: boolean) => void;
}

export function ViewDialog(props: ViewDialogProps) {
  const { children, view, open, onClose } = props;
  const { toast } = useToast();
  const [selectedResourceId, setSelectedResourceId] = useState<number | null>(
    null
  );
  const [id, setId] = useState<number | undefined>(undefined);
  const [name, setName] = useState('');
  const [ql, setQl] = useState('');
  const [option, setOption] = useState<Option>({
    columns: [],
    parameters: [],
    version: 0,
  });

  useEffect(() => {
    console.log('view', view);
    if (view) {
      setId(view.id);
      setName(view.name);
      setSelectedResourceId(view.resourceId);
      setQl(view.ql);
      setOption(JSON.parse(view.options || '{}'));
    } else {
      setId(undefined);
      setName('');
      setQl('');
      setSelectedResourceId(null);
      setOption({ columns: [], parameters: [], version: 0 });
    }
  }, [view]);

  const inputedSqlHandle = (s: string) => {
    if (s === ql) {
      return;
    }
    if (!s) {
      setQl('');
      setOption({ columns: [], parameters: [], version: 0 });
      return;
    }
    const pr = parse(s, 'MySQL');
    setQl(pr.sql || '');

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
    if (option.version != newOption.version) {
      setOption(newOption);
    }
  };

  const saveHandle = () => {
    const saveData: View = {
      id: id,
      name: name,
      resourceId: selectedResourceId || null,
      ql: ql,
      options: JSON.stringify(option),
    };
    if (!viewMetadata.OnValidate(saveData)) {
      return;
    }
    baseMutate(viewMetadata.Save, saveData)
      .then(() => onClose(true))
      .catch((e) => toast({ variant: 'destructive', title: e.message }));
  };

  return (
    <Dialog
      open={open}
      onOpenChange={(v) => {
        console.log('onOpenChange', v);
        if (!v) {
          setName('');
          setSelectedResourceId(null);
          setOption({ columns: [], parameters: [], version: 0 });
          onClose(false);
        }
      }}
    >
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-w-max'>
        <DialogHeader>
          <DialogTitle>View Editor</DialogTitle>
        </DialogHeader>
        {open && (
          <>
            <div className='grid grid-cols-8 items-center gap-4 w-[400px]'>
              <div className='col-span-2'>
                <Label>View name</Label>
              </div>
              <div className='col-span-6'>
                <Input
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                ></Input>
              </div>
            </div>
            <div className='gap-1 flex flex-row'>
              <div className=''>
                <ResourcePanel
                  onSelectedResource={setSelectedResourceId}
                ></ResourcePanel>
              </div>
              <div className=''>
                <SqlEditor ql={ql} onInputedQl={inputedSqlHandle}></SqlEditor>
              </div>
              <div className=''>
                <OptionPanel
                  option={option}
                  onEditedOption={editedOptionHandle}
                ></OptionPanel>
              </div>
            </div>
          </>
        )}
        <DialogFooter>
          <Button onClick={saveHandle}>Save</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
