import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';

import { cn } from '@/lib/utils';
import ResourceEditor from '@/components/editor/resource-editor';

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Resource } from '@/types';
import { MixerHorizontalIcon } from '@radix-ui/react-icons';
import { useEffect, useState } from 'react';
import { baseQuery } from '@/lib/graphql';

export default function ResourcePanel() {
  const [editing, setEditing] = useState(false);
  const [resources, setResources] = useState([] as Resource[]);
  const [resource, setResource] = useState(null as Resource | null);
  const [selected, setSelected] = useState(null as number | null);

  const onSelected = (resource: Resource) => {
    console.log('onSelected', resource.id);
    setSelected(resource.id ?? null);
  };

  const handleEdit = (resource: Resource) => {
    setResource(resource);
    setEditing(true);
  };
  const handleAdd = () => {
    console.log('handleAdd');
    setResource(null);
    setEditing(true);
  };

  const fetchResources = () => {
    console.log('fetchResources');
    baseQuery<{ resources: Resource[] }>(`{
        resources {
          id
          name
          type
          data
        }
      }`).then((data) => setResources(data.resources));
    setEditing(false);
  };

  useEffect(() => {
    fetchResources();
  }, []);

  return (
    <>
      <Card className='w-[250px] h-[400px]'>
        <CardHeader>
          <CardTitle>Resource</CardTitle>
        </CardHeader>
        <CardContent className='mb-0 pb-0'>
          <ScrollArea className='h-[280px]'>
            <div
              className={cn(
                'text-sm max-w-[200px]',
                resources.length > 6 ? 'pr-3' : ''
              )}
            >
              {resources.map((resource) => (
                <div
                  key={resource.id}
                  className={cn(
                    selected === resource.id
                      ? 'border-primary bg-primary text-background'
                      : 'border-background border-dashed hover:border-dashed hover:border-primary',
                    'mt-2 mb-2 border pt-1 pb-1 pl-2 pr-2 rounded-lg hover:cursor-pointer flex flex-row'
                  )}
                  onClick={() => onSelected(resource)}
                >
                  <div className='inline-block flex-1 truncate '>
                    {resource.name}
                  </div>
                  <MixerHorizontalIcon
                    className='w-5 h-5 inline-block hover:text-green-500 self-center '
                    onClick={(e) => handleEdit(resource)}
                  />
                </div>
              ))}
            </div>
          </ScrollArea>
        </CardContent>
        <CardFooter className='pt-2'>
          <Button
            variant='outline'
            className='w-full justify-center text-slate-400 border-slate-400 hover:text-slate-600 hover:border-slate-600 hover:bg-inherit border-2 border-dashed'
            onClick={handleAdd}
          >
            Add resource
            {editing ? '1' : '2'}
          </Button>
        </CardFooter>
      </Card>
      {editing ? (
        <ResourceEditor
          resource={resource}
          onRefresh={fetchResources}
        ></ResourceEditor>
      ) : null}
    </>
  );
}
