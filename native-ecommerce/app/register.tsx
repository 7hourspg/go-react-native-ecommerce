import { AuthHeader, AuthLayout, RegisterForm } from '@/components/auth';
import { Redirect } from 'expo-router';
import { useAuthState } from '@/context/auth-context';

export default function RegisterScreen() {
  return (
    <AuthLayout>
      <AuthHeader title="Create Account" subtitle="Sign up to start shopping" />
      <RegisterForm />
    </AuthLayout>
  );
}
