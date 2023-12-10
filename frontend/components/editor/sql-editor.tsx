import { Textarea } from '@/components/ui/textarea';
import { cn } from '@/lib/utils';
import { ChangeEvent, useState } from 'react';
import parse from '@/lib/sql';
import { useToast } from '../ui/use-toast';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
export interface SqlEditorProps {
  onInputedSql: (sql: string) => void;
}

export default function SqlEditor(props: SqlEditorProps) {
  const { onInputedSql } = props;
  const [timeoutState, setTimeoutState] = useState<NodeJS.Timeout | undefined>(
    undefined
  );
  const [error, setError] = useState<boolean>(false);
  const { toast } = useToast();
  const onChangeHandle = (event: ChangeEvent<HTMLTextAreaElement>) => {
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
