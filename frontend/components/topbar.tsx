import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Component1Icon, GitHubLogoIcon } from '@radix-ui/react-icons';
import { useTheme } from 'next-themes';

export default function Topbar() {
  const { theme, setTheme } = useTheme();

  return (
    <>
      <div className='absolute w-full flex justify-between items-center bg-secondary h-12'>
        <div className='pl-4 subpixel-antialiased font-semibold'>
          DataViewer
        </div>
        <div className='flex justify-between items-center'>
          <div className='mr-4 cursor-pointer'>
            <GitHubLogoIcon className='h-6 w-6' />
          </div>
          <div className='mr-4'>
            <Avatar className='h-6 w-6'>
              <AvatarImage src='https://github.com/shadcn.png' />
              <AvatarFallback>L</AvatarFallback>
            </Avatar>
          </div>
        </div>
      </div>
    </>
  );
}
