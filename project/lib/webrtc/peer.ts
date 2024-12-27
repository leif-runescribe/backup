export const createPeerConnection = () => {
  const configuration: RTCConfiguration = {
    iceServers: [
      { urls: 'stun:stun.l.google.com:19302' },
      { urls: 'stun:stun1.l.google.com:19302' },
    ],
  };

  return new RTCPeerConnection(configuration);
};

export const addTrack = (pc: RTCPeerConnection, stream: MediaStream) => {
  stream.getTracks().forEach((track) => {
    pc.addTrack(track, stream);
  });
};

export const createOffer = async (pc: RTCPeerConnection) => {
  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);
  return offer;
};

export const handleAnswer = async (pc: RTCPeerConnection, answer: RTCSessionDescriptionInit) => {
  await pc.setRemoteDescription(new RTCSessionDescription(answer));
};

export const handleIceCandidate = async (pc: RTCPeerConnection, candidate: RTCIceCandidateInit) => {
  await pc.addIceCandidate(new RTCIceCandidate(candidate));
};