export interface PeerConnection {
  connection: RTCPeerConnection;
  stream: MediaStream | null;
}

export interface User {
  id: string;
  username: string;
  points: number;
}

export interface RoomState {
  id: string;
  name: string;
  participants: User[];
}