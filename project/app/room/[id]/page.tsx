
import { VideoRoom } from '@/components/video-room';
import { useUserStore } from '@/lib/store';
import { redirect } from 'next/navigation';

// This is required for static site generation with dynamic routes
export function generateStaticParams() {
  return [];
}

export default function RoomPage({ params }: { params: { id: string } }) {
  const username = useUserStore((state) => state.username);

  if (!username) {
    redirect('/');
  }

  return (
    <div className="h-screen">
      <VideoRoom roomId={params.id} />
    </div>
  );
}