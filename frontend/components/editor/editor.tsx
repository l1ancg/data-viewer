import { MouseEvent, useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from '@/components/ui/dialog';
import { baseMutate } from '@/lib/graphql';
import { useToast } from '@/components/ui/use-toast';
import { Switch } from '@/components/ui/switch';

export interface MyOption {
  value: string;
  label: string;
}

// 声明字段类型，包含一个name和一个type
export interface MyField {
  name: string;
  type: string;
  options?: Array<MyOption>;
}

export interface MyEditorProps<T> {
  row: T | null;
  mutate: string;
  fields: MyField[];
  onRefresh?: () => void;
  onValidate?: (row: T) => void;
  title?: string;
  desc?: string;
}

export function MyEditor({
  row,
  mutate,
  fields,
  onRefresh,
  onValidate,
  title,
  desc,
}: MyEditorProps<any>) {
  const { toast } = useToast();
  const [data, setData] = useState(row);
  const handleChange = (name: string, value: string | boolean) => {
    setData({
      ...data,
      [name]: value,
    });
  };

  const handleSubmit = (event: MouseEvent<HTMLButtonElement>) => {
    console.log(111, data);
    event.preventDefault();
    try {
      onValidate && onValidate(data);
    } catch (e: any) {
      toast({ title: e.message });
      return;
    }

    baseMutate(mutate, data)
      .then(() => onRefresh && onRefresh())
      .catch((e) => console.log(e));
  };

  return (
    <>
      <Dialog defaultOpen onOpenChange={() => onRefresh && onRefresh()}>
        <DialogContent>
          <DialogHeader>
            {title && <DialogTitle>{title}</DialogTitle>}
            {desc && <DialogDescription>{desc}</DialogDescription>}
          </DialogHeader>
          <form className='grid gap-4 py-4'>
            {fields.map(
              (field) =>
                (field.type === 'input' && (
                  <div
                    className='grid grid-cols-5 items-center gap-4'
                    key={field.name}
                  >
                    <Label className='text-right'>{field.name}</Label>
                    <Input
                      className='col-span-4'
                      value={data[field.name]}
                      onChange={(e) => handleChange(field.name, e.target.value)}
                      required
                    />
                  </div>
                )) ||
                (field.type === 'select' && (
                  <div
                    className='grid grid-cols-5 items-center gap-4'
                    key={field.name}
                  >
                    <Label className='text-right'>{field.name}</Label>
                    <Select
                      value={data[field.name]}
                      onValueChange={(v) => handleChange(field.name, v)}
                      required
                    >
                      <SelectTrigger className='col-span-4'>
                        <SelectValue placeholder='' />
                      </SelectTrigger>
                      <SelectContent>
                        {field.options &&
                          field.options.map((option) => (
                            <SelectItem value={option.value} key={option.value}>
                              {option.label}
                            </SelectItem>
                          ))}
                      </SelectContent>
                    </Select>
                  </div>
                )) ||
                (field.type === 'textarea' && (
                  <div
                    className='grid grid-cols-5 items-center gap-4'
                    key={field.name}
                  >
                    <Label className='text-right'>{field.name}</Label>
                    <Textarea
                      className='col-span-4'
                      value={data[field.name]}
                      onChange={(e) => handleChange(field.name, e.target.value)}
                    />
                  </div>
                )) ||
                (field.type === 'switch' && (
                  <div
                    className='grid grid-cols-5 items-center gap-4'
                    key={field.name}
                  >
                    <Label className='text-right'>{field.name}</Label>
                    <Switch
                      className='col-span-4'
                      checked={data[field.name]}
                      onCheckedChange={(v) => handleChange(field.name, v)}
                    />
                  </div>
                ))
            )}

            <div className='gap-4 grid-cols-5 flex justify-end'>
              {/* <Button variant='destructive' onClick={(e) => handleDelete(e)}>
                Delete
              </Button> */}
              <Button type='submit' onClick={(e) => handleSubmit(e)}>
                Save
              </Button>
            </div>
          </form>
        </DialogContent>
      </Dialog>
    </>
  );
}
