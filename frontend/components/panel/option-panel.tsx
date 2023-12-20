import { useEffect, useState } from 'react';
import { useToast } from '@/components/ui/use-toast';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Option } from '@/types';
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import { ComponentInstanceIcon } from '@radix-ui/react-icons';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';

export interface OptionPanelProps {
  option: Option;
  onEditedOption: (option: Option) => void;
}

export default function OptionPanel(props: OptionPanelProps) {
  const { option } = props;
  const { toast } = useToast();
  const [localOption, setLocalOption] = useState(props.option);

  useEffect(() => {
    setLocalOption(props.option);
  }, [props.option]);

  useEffect(() => {
    props.onEditedOption(localOption);
  }, [localOption]);

  const columnLabelChangeHandle = (idx: number, val: string) => {
    setLocalOption((prev) => {
      const next = { ...prev };
      next.columns[idx].label = val;
      next.version = next.version + 1;
      return next;
    });
  };
  const parameterLabelChangeHandle = (idx: number, val: string) => {
    setLocalOption((prev) => {
      const next = { ...prev };
      next.parameters[idx].label = val;
      next.version = next.version + 1;
      return next;
    });
  };

  return (
    <>
      <Card className='w-[250px] h-[500px]'>
        <CardHeader>
          <CardTitle>Option</CardTitle>
        </CardHeader>
        <CardContent className='mb-0 pb-0'>
          <ScrollArea className='h-[380px]'>
            <div className='m-2'>
              <ComponentInstanceIcon className='inline mr-2' />
              <span>Column</span>
            </div>
            {option.columns.length > 0 ? (
              <Accordion type='single' collapsible className='w-full'>
                {option.columns.map((column, idx) => (
                  <AccordionItem value='item-2' key={column.name}>
                    <AccordionTrigger>{column.name}</AccordionTrigger>
                    <AccordionContent>
                      <div className='grid grid-cols-4 items-center gap-4'>
                        <Label htmlFor='name' className='text-right'>
                          Label
                        </Label>
                        <Input
                          id='name'
                          className='col-span-3 focus-visible:ring-0'
                          value={column.label}
                          onChange={(e) => {
                            columnLabelChangeHandle(idx, e.target.value);
                          }}
                        />
                      </div>
                    </AccordionContent>
                  </AccordionItem>
                ))}
              </Accordion>
            ) : (
              <div className='text-center text-gray-400'>empty</div>
            )}
            <div className='m-2'>
              <ComponentInstanceIcon className='inline mr-2' />
              <span>Paramter</span>
            </div>
            {option.parameters.length > 0 ? (
              <Accordion type='single' collapsible className='w-full'>
                {option.parameters.map((param, idx) => (
                  <AccordionItem value='item-2' key={param.name}>
                    <AccordionTrigger>{param.name}</AccordionTrigger>
                    <AccordionContent>
                      <div className='grid grid-cols-4 items-center gap-4'>
                        <Label htmlFor='name' className='text-right'>
                          Label
                        </Label>
                        <Input
                          id='name'
                          className='col-span-3 focus-visible:ring-0'
                          value={param.label}
                          onChange={(e) => {
                            parameterLabelChangeHandle(idx, e.target.value);
                          }}
                        />
                      </div>
                    </AccordionContent>
                  </AccordionItem>
                ))}
              </Accordion>
            ) : (
              <div className='text-center text-gray-400'>empty</div>
            )}
          </ScrollArea>
        </CardContent>
      </Card>
    </>
  );
}
