'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { RoomCard } from '@/components/room-card';
import { Plus } from 'lucide-react';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { v4 as uuidv4 } from 'uuid';

interface Room {
  id: string;
  name: string;
  participants: number;
}

export default function RoomsPage() {
  const [rooms, setRooms] = useState<Room[]>([
    { id: '1', name: 'Gaming Lounge', participants: 3 },
    { id: '2', name: 'Study Group', participants: 2 },
  ]);
  const [newRoomName, setNewRoomName] = useState('');
  const router = useRouter();

  const createRoom = () => {
    if (newRoomName.trim()) {
      const newRoom = {
        id: uuidv4(),
        name: newRoomName.trim(),
        participants: 0,
      };
      setRooms([...rooms, newRoom]);
      setNewRoomName('');
      router.push(`/room/${newRoom.id}`);
    }
  };

  const joinRoom = (roomId: string) => {
    router.push(`/room/${roomId}`);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold">Available Rooms</h1>
        <Dialog>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> Create Room
            </Button>
          </DialogTrigger>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Create New Room</DialogTitle>
            </DialogHeader>
            <div className="space-y-4 pt-4">
              <Input
                placeholder="Room name"
                value={newRoomName}
                onChange={(e) => setNewRoomName(e.target.value)}
              />
              <Button onClick={createRoom} className="w-full">
                Create
              </Button>
            </div>
          </DialogContent>
        </Dialog>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {rooms.map((room) => (
          <RoomCard
            key={room.id}
            {...room}
            onJoin={() => joinRoom(room.id)}
          />
        ))}
      </div>
    </div>
  );
}