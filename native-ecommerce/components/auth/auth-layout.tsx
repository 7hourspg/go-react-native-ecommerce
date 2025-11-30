import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ReactNode } from 'react';
import { useAuthState } from '@/context/auth-context';
import { Redirect } from 'expo-router';

interface AuthLayoutProps {
  children: ReactNode;
}

export function AuthLayout({ children }: AuthLayoutProps) {
  const { isLoggedIn } = useAuthState();

  // Redirect if already logged in
  if (isLoggedIn) {
    return <Redirect href="/(tabs)" />;
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top', 'bottom']}>
      <ScrollView
        className="flex-1"
        contentContainerStyle={{ flexGrow: 1 }}
        showsVerticalScrollIndicator={false}>
        <View className="flex-1 justify-center px-6 py-8">{children}</View>
      </ScrollView>
    </SafeAreaView>
  );
}
