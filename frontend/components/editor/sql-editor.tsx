import { Textarea } from '@/components/ui/textarea';
import { cn } from '@/lib/utils';
import { ChangeEvent, useEffect, useState } from 'react';
import parse from '@/lib/sql';
import { useToast } from '../ui/use-toast';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
export interface SqlEditorProps {
  ql: string;
  onInputedQl: (sql: string) => void;
}

export default function SqlEditor(props: SqlEditorProps) {
  const { onInputedQl: onInputedSql, ql } = props;
  const [timeoutState, setTimeoutState] = useState<NodeJS.Timeout | undefined>(
    undefined
  );
  const [textAreaValue, setTextAreaValue] = useState<string>('');
  const [error, setError] = useState<boolean>(false);
  const { toast } = useToast();

  useEffect(() => {
    setTextAreaValue(ql);
  }, [ql]);

  const onChangeHandle = (event: ChangeEvent<HTMLTextAreaElement>) => {
    setTextAreaValue(event.target.value);
    delay(event.target.value);
  };
  const delay = (s: string) => {
    if (!s) {
      setError(false);
      onInputedSql('');
      return;
    }
    clearTimeout(timeoutState);
    setTimeoutState(
      setTimeout(() => {
        onInputedSql(s);
        const { error } = parse(s, 'MySQL');
        if (error) {
          toast({
            variant: 'destructive',
            title:
              error.start.line === error.end.line
                ? `Syntax error: line ${error.start.line}`
                : `Syntax error: line ${error.start.line} - ${error.end.line}`,
          });
          setError(true);
        } else {
          setError(false);
        }
      }, 1000)
    );
  };
  return (
    <>
      <Card className='w-[500px] h-[500px]'>
        <CardHeader>
          <CardTitle>Sql</CardTitle>
        </CardHeader>
        <CardContent className='mb-0 pb-0'>
          <Textarea
            className={cn(
              'focus-visible:ring-0',
              error ? 'border-red-500' : ''
            )}
            value={textAreaValue}
            style={{ resize: 'none' }}
            cols={70}
            rows={20}
            placeholder='example: select name, age from users where age > 18'
            onChange={onChangeHandle}
          ></Textarea>
        </CardContent>
      </Card>
    </>
  );
}
