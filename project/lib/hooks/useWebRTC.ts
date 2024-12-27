'use client';

import { useState, useCallback, useEffect } from 'react';
import { getLocalStream, toggleTrack } from '@/lib/webrtc/media';
import { createPeerConnection, addTrack } from '@/lib/webrtc/peer';
import type { PeerConnection } from '@/lib/webrtc/types';
import { io, Socket } from 'socket.io-client';

export function useWebRTC(roomId: string, username: string) {
  const [localStream, setLocalStream] = useState<MediaStream | null>(null);
  const [peers, setPeers] = useState<Map<string, PeerConnection>>(new Map());
  const [videoEnabled, setVideoEnabled] = useState(true);
  const [audioEnabled, setAudioEnabled] = useState(true);
  const [socket, setSocket] = useState<Socket | null>(null);

  const initializeMedia = useCallback(async () => {
    try {
      const stream = await getLocalStream();
      setLocalStream(stream);
    } catch (error) {
      console.error('Failed to get local stream:', error);
    }
  }, []);

  useEffect(() => {
    initializeMedia();
    const newSocket = io(process.env.NEXT_PUBLIC_SIGNALING_SERVER || 'http://localhost:3001');
    setSocket(newSocket);

    newSocket.on('connect', () => {
      newSocket.emit('join-room', { roomId, username });
    });

    newSocket.on('user-joined', async ({ userId, username }) => {
      if (localStream) {
        const pc = createPeerConnection();
        addTrack(pc, localStream);
        setPeers((prev) => new Map(prev.set(userId, { connection: pc, stream: null })));
      }
    });

    return () => {
      localStream?.getTracks().forEach((track) => track.stop());
      peers.forEach((peer) => peer.connection.close());
      newSocket.close();
    };
  }, [roomId, username, localStream]);

  const handleToggleVideo = useCallback(() => {
    if (localStream) {
      toggleTrack(localStream, 'video', !videoEnabled);
      setVideoEnabled(!videoEnabled);
    }
  }, [localStream, videoEnabled]);

  const handleToggleAudio = useCallback(() => {
    if (localStream) {
      toggleTrack(localStream, 'audio', !audioEnabled);
      setAudioEnabled(!audioEnabled);
    }
  }, [localStream, audioEnabled]);

  return {
    localStream,
    peers,
    videoEnabled,
    audioEnabled,
    handleToggleVideo,
    handleToggleAudio,
  };
}