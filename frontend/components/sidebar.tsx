import { cn } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { ViewDialog } from '@/components/view-dialog';
import { baseQuery } from '@/lib/graphql';
import { View } from '@/types';

interface SidebarProps extends React.HTMLAttributes<HTMLDivElement> {}
const dashboard: View = {
  id: 0,
  resourceId: 1,
  resourceType: 'dashboard',
  displayType: '',
  name: 'Dashboard',
  desc: '',
};
export default function Sidebar({ className }: SidebarProps) {
  const [views, setViews] = useState<View[]>([dashboard]);

  const [active, setActive] = useState<number>(0);
  const fetchViews = () => {
    return baseQuery<{ views: View[] }>(`{
        views {
          id
          resourceId
          resourceType
          displayType
          name
          desc
        }
      }`);
  };

  useEffect(() => {
    fetchViews().then((data) => {
      setViews([dashboard, ...data.views]);
    });
  }, []);

  const router = useRouter();
  const onActive = (item: View) => {
    if (item.resourceType === 'dashboard') {
      router.push(`/`);
    } else if (item.resourceType === 'mysql') {
      router.push(`/mysql/${item.id}`);
    }
    setActive(item.id);
  };

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
                  variant={active === item.id ? 'default' : 'ghost'}
                  className='w-full justify-start'
                  key={item.id.toString()}
                  onClick={() => onActive(item)}
                >
                  {item.name}
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
