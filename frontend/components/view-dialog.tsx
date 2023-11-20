import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import ResourcePanel from '@/components/panel/resource-panel';
import ColumnPanel from '@/components/panel/column-panel';
import { useState, ReactNode } from 'react';

export function ViewDialog({ children }: { children: ReactNode }) {
  const [selectedResourceId, setSelectedResourceId] = useState<number | null>(
    null
  );
  const selectedResourceHandle = (id: number | null) => {
    setSelectedResourceId(id);
  };

  const openChangeHandle = (open: boolean) => {
    selectedResourceHandle(null);
  };

  return (
    <Dialog onOpenChange={openChangeHandle}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-w-max'>
        <DialogHeader>
          <DialogTitle>View Editor</DialogTitle>
        </DialogHeader>
        <div className='gap-3 flex flex-row'>
          <div className=''>
            <ResourcePanel onSelected={selectedResourceHandle}></ResourcePanel>
          </div>
          <div className=''>
            <ColumnPanel resourceId={selectedResourceId}></ColumnPanel>
          </div>
        </div>
        <DialogFooter>
          <Button type='submit'>Save</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
