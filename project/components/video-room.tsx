'use client';

import { LocalVideo } from '@/components/video/local-video';
import { RemoteVideo } from '@/components/video/remote-video';
import { VideoControls } from '@/components/video/controls';
import { useWebRTC } from '@/lib/hooks/useWebRTC';
import { useUserStore } from '@/lib/store';

interface VideoRoomProps {
  roomId: string;
}

export function VideoRoom({ roomId }: VideoRoomProps) {
  const username = useUserStore((state) => state.username);
  const {
    localStream,
    peers,
    videoEnabled,
    audioEnabled,
    handleToggleVideo,
    handleToggleAudio,
  } = useWebRTC(roomId, username);

  return (
    <div className="h-full flex flex-col">
      <div className="flex-1 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
        <LocalVideo
          stream={localStream}
          username={username}
          className="w-full aspect-video"
        />
        {Array.from(peers.entries()).map(([peerId, peer]) => (
          peer.stream && (
            <RemoteVideo
              key={peerId}
              stream={peer.stream}
              username={`User ${peerId}`}
              className="w-full aspect-video"
            />
          )
        ))}
      </div>
      <VideoControls
        videoEnabled={videoEnabled}
        audioEnabled={audioEnabled}
        onToggleVideo={handleToggleVideo}
        onToggleAudio={handleToggleAudio}
      />
    </div>
  );
}