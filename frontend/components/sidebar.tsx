import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { ViewDialog } from '@/components/view-dialog';
import { View } from '@/types';
import { baseQuery, baseMutate } from '@/lib/graphql';
import viewMetadata from '@/components/metadata/view';
import { useToast } from '@/components/ui/use-toast';
import viewStore from '@/lib/viewStore';
import {
  ContextMenu,
  ContextMenuTrigger,
  ContextMenuContent,
  ContextMenuItem,
} from '@/components/ui/context-menu';

interface SidebarProps extends React.HTMLAttributes<HTMLDivElement> {}
const dashboard: View = {
  id: -1,
  resourceId: -1,
  name: 'Dashboard',
  ql: '',
  options: '',
};
export default function Sidebar({ className }: SidebarProps) {
  const { toast } = useToast();
  const [views, setViews] = useState<View[]>([dashboard]);
  const [active, setActive] = useState<number>(0);
  const [editingView, setEditingView] = useState<View | null>(null);
  const [openEditing, setOpenEditing] = useState(false);
  // todo first delete, second delete

  const fetchViews = () => {
    return baseQuery<{ views: View[] }>(viewMetadata.Query)
      .then((data) => setViews([dashboard, ...data.views]))
      .catch((e) => toast({ variant: 'destructive', title: e.message }));
  };

  useEffect(() => {
    fetchViews();
  }, []);

  const router = useRouter();
  const onActive = (item: View) => {
    if (item.id === -1) {
      router.push(`/`);
      setActive(-1);
      return;
    }
    viewStore.setView(item);
    router.push(`/mysql`);
    setActive(item.id || -1);
  };

  const onDelete = (item: View) => {
    baseMutate(viewMetadata.Delete, { id: item.id }).then(() => fetchViews());
  };

  return (
    <div className={cn('pb-2', className)}>
      <div className='pt-2 m-1 bg-secondary	rounded-lg h-full flex flex-col justify-between'>
        <div className='px-3 py-2'>
          {/* <h2 className='mb-2 px-2 text-slate-400 font-semibold'>Discover</h2> */}
          <div className='space-y-1'>
            {/* secondary ghost */}
            {views.length > 0 ? (
              views.map((item: View) => (
                <ContextMenu key={item.id?.toString()}>
                  <ContextMenuTrigger>
                    <Button
                      variant={active === item.id ? 'default' : 'ghost'}
                      className={cn(
                        'w-full justify-start my-1',
                        active === item.id ? '' : 'hover:bg-slate-200'
                      )}
                      onClick={() => onActive(item)}
                    >
                      {item.name}
                    </Button>
                  </ContextMenuTrigger>
                  <ContextMenuContent>
                    <ContextMenuItem
                      onClick={() => {
                        setEditingView(item);
                        setOpenEditing(true);
                      }}
                    >
                      Properties
                    </ContextMenuItem>
                    <ContextMenuItem
                      className='focus:bg-red-100'
                      onClick={() => {
                        onDelete(item);
                      }}
                    >
                      Delete
                    </ContextMenuItem>
                  </ContextMenuContent>
                </ContextMenu>
              ))
            ) : (
              <div className='text-slate-400 text-center text-sm'>No view</div>
            )}
          </div>
        </div>

        <div className='pt-2 m-1 px-3 py-2'>
          <ViewDialog
            open={openEditing}
            view={editingView}
            onClose={(refresh) => {
              setEditingView(null);
              setOpenEditing(false);
              if (refresh) {
                fetchViews();
              }
            }}
          >
            <Button
              variant='outline'
              className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
              onClick={() => {
                setEditingView(null);
                setOpenEditing(true);
              }}
            >
              Add
            </Button>
          </ViewDialog>
          <div className='mb-1 mt-2 text-center text-sm text-muted-foreground'>
            Version: 1.0.0
          </div>
        </div>
      </div>
    </div>
  );
}
