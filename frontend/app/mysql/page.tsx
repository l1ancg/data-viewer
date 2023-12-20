'use client';
import { View, Option } from '@/types';
import viewStore from '@/lib/viewStore';
import { observer } from 'mobx-react';
import { useState, useEffect, use } from 'react';
import api from '@/lib/api';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import TablePanel from '@/components/panel/table-panel';

export default observer(function Page() {
  const [view, setView] = useState<View | null>(null);
  const [option, setOption] = useState<Option | null>(null);
  const [parameter, setParameter] = useState<any>({});
  const [data, setData] = useState<any>([]);

  useEffect(() => {
    setView(viewStore.view);
    setOption(JSON.parse(viewStore.view?.options || '{}'));
  }, [viewStore.view]);

  useEffect(() => {
    if (view) {
      fetchData();
    }
  }, [view]);

  const fetchData = () => {
    console.log('before fetchData', view);
    api
      .post('/query', {
        resourceId: view?.resourceId,
        ql: view?.ql,
        parameter: parameter,
      })
      .then((res) => {
        setData(res);
      });
  };

  const queryHandle = () => {
    if (!view) {
      return;
    }
    fetchData();
  };

  const changedParameterHandle = (name: string, val: string) => {
    setParameter({ ...parameter, [name]: val });
  };

  return (
    <>
      <div>View Name: {view?.name}</div>
      {option?.parameters?.length || 0 > 0 ? (
        <div className='flex items-center justify-between'>
          <div className='flex flex-1 items-center space-x-2'>
            {option?.parameters.map((item) => (
              <Input
                key={item.name}
                placeholder={'Filter ' + item.label + '...'}
                onChange={(event) =>
                  changedParameterHandle(item.name, event.target.value)
                }
                className='h-8 w-[150px] lg:w-[250px] focus-visible:ring-0'
              />
            ))}
            <Button onClick={queryHandle} className='h-8 px-2 lg:px-3'>
              Query
            </Button>
          </div>
        </div>
      ) : null}
      <TablePanel data={data} columns={option?.columns || []}></TablePanel>
    </>
  );
});
