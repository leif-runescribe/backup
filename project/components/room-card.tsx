'use client';

import { Card, CardHeader, CardTitle, CardContent, CardFooter } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Users } from 'lucide-react';

interface RoomCardProps {
  id: string;
  name: string;
  participants: number;
  onJoin: () => void;
}

export function RoomCard({ id, name, participants, onJoin }: RoomCardProps) {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="text-xl font-bold">{name}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="flex items-center gap-2">
          <Users className="h-4 w-4" />
          <span>{participants} participants</span>
        </div>
      </CardContent>
      <CardFooter>
        <Button onClick={onJoin} className="w-full">
          Join Room
        </Button>
      </CardFooter>
    </Card>
  );
}