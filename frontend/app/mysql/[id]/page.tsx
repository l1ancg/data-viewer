export default function Page({ params }: { params: { id: string } }) {
  return <div>My ID: {params.id}</div>;
}
