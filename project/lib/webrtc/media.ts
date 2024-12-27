export const getLocalStream = async () => {
  try {
    return await navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true,
    });
  } catch (error) {
    console.error('Error accessing media devices:', error);
    throw error;
  }
};

export const toggleTrack = (stream: MediaStream, kind: 'audio' | 'video', enabled: boolean) => {
  stream.getTracks().forEach((track) => {
    if (track.kind === kind) {
      track.enabled = enabled;
    }
  });
};