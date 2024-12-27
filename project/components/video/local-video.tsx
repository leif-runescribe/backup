'use client';

import { useEffect, useRef } from 'react';
import { cn } from '@/lib/utils';

interface LocalVideoProps {
  stream: MediaStream | null;
  username: string;
  className?: string;
}

export function LocalVideo({ stream, username, className }: LocalVideoProps) {
  const videoRef = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    if (videoRef.current && stream) {
      videoRef.current.srcObject = stream;
    }
  }, [stream]);

  return (
    <div className={cn('relative rounded-lg overflow-hidden bg-gray-900', className)}>
      <video
        ref={videoRef}
        autoPlay
        playsInline
        muted
        className="h-full w-full object-cover"
      />
      <div className="absolute bottom-2 left-2 bg-black/50 px-2 py-1 rounded text-white">
        {username} (You)
      </div>
    </div>
  );
}