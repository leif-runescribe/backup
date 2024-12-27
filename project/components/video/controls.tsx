'use client';

import { Button } from '@/components/ui/button';
import { Video, VideoOff, Mic, MicOff } from 'lucide-react';

interface VideoControlsProps {
  videoEnabled: boolean;
  audioEnabled: boolean;
  onToggleVideo: () => void;
  onToggleAudio: () => void;
}

export function VideoControls({
  videoEnabled,
  audioEnabled,
  onToggleVideo,
  onToggleAudio,
}: VideoControlsProps) {
  return (
    <div className="p-4 bg-gray-100 dark:bg-gray-800 flex justify-center gap-4">
      <Button
        variant={videoEnabled ? 'default' : 'destructive'}
        size="icon"
        onClick={onToggleVideo}
      >
        {videoEnabled ? <Video className="h-4 w-4" /> : <VideoOff className="h-4 w-4" />}
      </Button>
      <Button
        variant={audioEnabled ? 'default' : 'destructive'}
        size="icon"
        onClick={onToggleAudio}
      >
        {audioEnabled ? <Mic className="h-4 w-4" /> : <MicOff className="h-4 w-4" />}
      </Button>
    </div>
  );
}