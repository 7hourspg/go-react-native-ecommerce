import '@/global.css';

import * as React from 'react';
import { NAV_THEME } from '@/lib/theme';
import { ThemeProvider } from '@react-navigation/native';
import { PortalHost } from '@rn-primitives/portal';
import { Stack } from 'expo-router';
import { StatusBar } from 'expo-status-bar';
import { useColorScheme } from 'nativewind';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { AuthProvider, useAuthState } from '@/context/auth-context';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { useOnlineManager } from '@/hooks/useOnlineManager';
import { useAppState } from '@/hooks/useAppState';
import { AppStateStatus, Platform } from 'react-native';
import { focusManager } from '@tanstack/react-query';
import { client } from '@/client/client.gen';
import axios from 'axios';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { PaymentProvider } from '@/components/custom/payment-provider';

export {
  // Catch any errors thrown by the Layout component.
  ErrorBoundary,
} from 'expo-router';

function RootLayoutNav() {
  const { colorScheme } = useColorScheme();

  const BASE = 'http://192.168.29.171:8080';

  const { tokens, setTokens, isLoggedIn } = useAuthState();

  const isRefreshing = React.useRef(false);

  client.setConfig({
    baseURL: BASE,
    headers: {
      Authorization: `Bearer ${tokens?.access_token}`,
    },
  });

  //  INTERCEPTOR TO REFRESH TOKEN
  client.instance.interceptors.response.use(undefined, async (error) => {
    // get the original request
    const originalRequest = error.config;

    // Check if it's a 401 error and we haven't already tried to refresh
    if (error.response?.status === 401 && !originalRequest._retry) {
      // Prevent multiple simultaneous refresh attempts
      if (isRefreshing.current) {
        return Promise.reject(error);
      }

      originalRequest._retry = true;
      isRefreshing.current = true;

      try {
        const tokensData = await AsyncStorage.getItem('@token_user');
        if (!tokensData) {
          return;
        }
        const refreshToken = JSON.parse(tokensData).refresh_token;

        const { data } = await axios.get(`${BASE}/auth/refresh`, {
          headers: {
            Authorization: `Bearer ${refreshToken}`,
          },
        });

        // Update tokens in storage and state
        await AsyncStorage.setItem('@token_user', JSON.stringify(data.tokens));
        setTokens(data.tokens);

        // Update the failed request's authorization header
        originalRequest.headers.Authorization = `Bearer ${data.tokens.access_token}`;

        isRefreshing.current = false;

        // Retry the original request
        return client.instance(originalRequest);
      } catch (refreshError) {
        isRefreshing.current = false;

        // If refresh fails, logout the user
        await AsyncStorage.removeItem('@token_user');
        await AsyncStorage.removeItem('@user_user');
        setTokens(null);

        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  });

  return (
    <ThemeProvider value={NAV_THEME[colorScheme ?? 'light']}>
      <StatusBar style={colorScheme === 'dark' ? 'light' : 'dark'} />
      <Stack screenOptions={{ headerShown: false }}>
        {/* Public Auth Routes */}
        <Stack.Screen name="index" />
        <Stack.Screen name="sign-in" />
        <Stack.Screen name="register" />
        <Stack.Screen name="forgot-password" />

        {/* Protected Tabs Route */}
        <Stack.Protected guard={isLoggedIn}>
          <Stack.Screen name="(tabs)" />
        </Stack.Protected>
      </Stack>
      <PortalHost />
    </ThemeProvider>
  );
}

function onAppStateChange(status: AppStateStatus) {
  // React Query already supports in web browser refetch on window focus by default
  if (Platform.OS !== 'web') {
    focusManager.setFocused(status === 'active');
  }
}

export default function RootLayout() {
  useOnlineManager();

  useAppState(onAppStateChange);

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: 2,
        staleTime: 5 * 60 * 1000, // 5 minutes - data stays fresh for 5 minutes
        gcTime: 10 * 60 * 1000, // 10 minutes - cache persists for 10 minutes
      },
    },
  });

  return (
    <SafeAreaProvider>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <PaymentProvider> 
          <RootLayoutNav />
          </PaymentProvider>
        </AuthProvider>
      </QueryClientProvider>
    </SafeAreaProvider>
  );
}
