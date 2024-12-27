import { create } from 'zustand';

interface UserState {
  username: string;
  setUsername: (username: string) => void;
  points: number;
  addPoints: (points: number) => void;
}

export const useUserStore = create<UserState>((set) => ({
  username: '',
  setUsername: (username) => set({ username }),
  points: 0,
  addPoints: (points) => set((state) => ({ points: state.points + points })),
}));