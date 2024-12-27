'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { useUserStore } from '@/lib/store';
import { useRouter } from 'next/navigation';

export function UsernameForm() {
  const [inputUsername, setInputUsername] = useState('');
  const setUsername = useUserStore((state) => state.setUsername);
  const router = useRouter();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (inputUsername.trim()) {
      setUsername(inputUsername.trim());
      router.push('/rooms');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <Input
          type="text"
          placeholder="Enter your username"
          value={inputUsername}
          onChange={(e) => setInputUsername(e.target.value)}
          className="w-full"
          required
          minLength={3}
        />
      </div>
      <Button type="submit" className="w-full">
        Continue
      </Button>
    </form>
  );
}