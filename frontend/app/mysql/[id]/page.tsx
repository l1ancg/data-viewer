import { DataTable } from '@/components/data-table/data-table';
import { columns } from '@/components/columns';
export default function Page({ params }: { params: { id: string } }) {
  let tasks = [];

  for (let i = 0; i < 20; i++) {
    tasks.push({
      id: 'TASK-' + i,
      title: 'title' + i,
      status: 'in progress',
      label: 'documentation',
      priority: 'medium',
    });
  }
  return (
    <>
      {/* <div>My ID: {params.id}</div> */}
      <DataTable data={tasks} columns={columns} />
    </>
  );
}
