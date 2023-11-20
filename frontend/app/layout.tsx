'use client';
import './globals.css';
import Topbar from '@/components/topbar';
import Sidebar from '@/components/sidebar';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Toaster } from '@/components/ui/toaster';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang='en'>
      <body>
        <div className='h-screen w-screen min-h-page min-w-page font-mono'>
          <Topbar />

          <div className='pt-12 h-full w-full flex flex-row'>
            <Sidebar className='w-56 h-full' />
            <div className='col-span-3 h-full w-full'>
              <ScrollArea className='h-full px-4 py-6'>{children}</ScrollArea>
            </div>
          </div>
        </div>
        <Toaster />
      </body>
    </html>
  );
}
