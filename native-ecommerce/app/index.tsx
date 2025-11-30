import { Redirect } from 'expo-router';
import { useAuthState } from '@/context/auth-context';
import { Loading } from '@/components/custom/loading';

export default function Index() {
  const { isLoggedIn, isLoading } = useAuthState();

  if (isLoading) {
    return <Loading message="Loading..." />;
  }

  if (isLoggedIn) {
    return <Redirect href="/(tabs)/(home)" />;
  }

  return <Redirect href="/sign-in" />;
}
