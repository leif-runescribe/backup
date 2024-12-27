import { UsernameForm } from '@/components/username-form';

export default function Home() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-purple-500 to-pink-500">
      <div className="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-lg shadow-xl">
        <h1 className="text-3xl font-bold text-center mb-8">Welcome to GameMeet</h1>
        <p className="text-center text-gray-600 dark:text-gray-300 mb-8">
          Enter your username to start connecting with others
        </p>
        <UsernameForm />
      </div>
    </div>
  );
}