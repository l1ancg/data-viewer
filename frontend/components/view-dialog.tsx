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
import ColumnPanel from '@/components/panel/column-panel';

export function ViewDialog({ children }: { children: React.ReactNode }) {
  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className='max-w-max'>
        <DialogHeader>
          <DialogTitle>Add view</DialogTitle>
        </DialogHeader>
        <div className='gap-3 flex flex-row'>
          <div className=''>
            <ResourcePanel></ResourcePanel>
          </div>
          <div className=''>
            <ColumnPanel></ColumnPanel>
          </div>
        </div>
        <DialogFooter>
          <Button type='submit'>Save</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
