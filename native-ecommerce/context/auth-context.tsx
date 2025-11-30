import AsyncStorage from '@react-native-async-storage/async-storage';
import { createContext, useContext, useEffect, useState, ReactNode } from 'react';
import { ModelsUser, ModelsLoginResponse } from '@/client/types.gen';

interface AuthState {
  isLoggedIn: boolean;
  isLoading: boolean;
  tokens: { access_token: string; refresh_token: string } | null;
  login: (response: ModelsLoginResponse) => Promise<void>;
  user: ModelsUser | null;
  logout: () => Promise<void>;
  setTokens: (tokens: { access_token: string; refresh_token: string } | null) => void;
}

const AuthContext = createContext<AuthState | undefined>(undefined);

const TOKEN_STORAGE_KEY = '@token_user';
const USER_STORAGE_KEY = '@user_user';

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [tokens, setTokens] = useState<{ access_token: string; refresh_token: string } | null>(
    null
  );
  const [user, setUser] = useState<ModelsUser | null>(null);

  useEffect(() => {
    // Check if user is logged in on app start
    checkAuthState();
  }, []);

  const checkAuthState = async () => {
    try {
      const tokensData = await AsyncStorage.getItem(TOKEN_STORAGE_KEY);
      const userData = await AsyncStorage.getItem(USER_STORAGE_KEY);
      if (tokensData && userData) {
        const parsedTokens = JSON.parse(tokensData) as {
          access_token: string;
          refresh_token: string;
        };
        const parsedUser = JSON.parse(userData) as ModelsUser;
        setTokens({
          access_token: parsedTokens.access_token,
          refresh_token: parsedTokens.refresh_token,
        });
        setUser(parsedUser);
        setIsLoggedIn(true);
      }
    } catch (error) {
      console.error('Error checking auth state:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const login = async (response: ModelsLoginResponse) => {
    try {
      await AsyncStorage.setItem(TOKEN_STORAGE_KEY, JSON.stringify(response.tokens));
      await AsyncStorage.setItem(USER_STORAGE_KEY, JSON.stringify(response.user));
      setUser(response.user as ModelsUser);
      setTokens(response.tokens as { access_token: string; refresh_token: string });
      setIsLoggedIn(true);
    } catch (error) {
      console.error('Error logging in:', error);
      throw error;
    }
  };

  const logout = async () => {
    try {
      await AsyncStorage.removeItem(TOKEN_STORAGE_KEY);
      await AsyncStorage.removeItem(USER_STORAGE_KEY);
      setTokens(null);
      setUser(null);
      setIsLoggedIn(false);
    } catch (error) {
      console.error('Error logging out:', error);
      throw error;
    }
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, isLoading, tokens, login, logout, user, setTokens }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuthState() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuthState must be used within an AuthProvider');
  }
  return context;
}
