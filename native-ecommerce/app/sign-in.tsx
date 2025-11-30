import { AuthHeader, AuthLayout, SignInForm } from '@/components/auth';

export default function SignInScreen() {
  return (
    <AuthLayout>
      <AuthHeader title="Welcome Back" subtitle="Sign in to continue shopping" />
      <SignInForm />
    </AuthLayout>
  );
}
