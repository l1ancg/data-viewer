'use client';
import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { ViewDialog } from '@/components/view-dialog';

export interface View {
  name: string;
  id: Number;
  type: Number;
  visible?: boolean;
}

interface SidebarProps extends React.HTMLAttributes<HTMLDivElement> {}

export default function Sidebar({ className }: SidebarProps) {
  const [views, setViews] = useState<View[]>([
    // { name: '1', id: 1, type: 1, visible: true },
    // { name: '2', id: 2, type: 1, visible: false },
    // { name: '3', id: 3, type: 1, visible: false },
  ]);

  const router = useRouter();

  function onOpen(item: View) {
    const updatedViews = views.map((view) => ({
      ...view,
      visible: view.id === item.id,
    }));
    setViews(updatedViews);
    if (item.type === 0) {
      router.push(`/`);
    } else if (item.type === 1) {
      router.push(`/mysql/${item.id}`);
    }
  }

  function onAdd() {}

  return (
    <div className={cn('pb-2', className)}>
      <div className='pt-2 m-1 bg-secondary	rounded-lg h-full flex flex-col justify-between'>
        <div className='px-3 py-2'>
          {/* <h2 className='mb-2 px-2 text-slate-400 font-semibold'>Discover</h2> */}
          <div className='space-y-1'>
            {/* secondary ghost */}

            {views.length > 0 ? (
              views.map((item: View) => (
                <Button
                  variant={item.visible ? 'secondary' : 'ghost'}
                  className='w-full justify-start'
                  key={item.name}
                  onClick={() => onOpen(item)}
                >
                  示例🧭{item.name}
                </Button>
              ))
            ) : (
              <div className='text-slate-400 text-center text-sm'>No view</div>
            )}
          </div>
        </div>

        <div className='pt-2 m-1 px-3 py-2'>
          <ViewDialog>
            <Button
              variant='outline'
              className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
              onClick={() => onAdd()}
            >
              Add view
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
