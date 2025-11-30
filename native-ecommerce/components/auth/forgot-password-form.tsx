import { Button } from '@/components/ui/button';
import { IconInput } from './icon-input';
import { AuthCard } from './auth-card';
import { MailIcon } from 'lucide-react-native';
import { Link, useRouter } from 'expo-router';
import { Text } from '@/components/ui/text';
import { View } from 'react-native';
import { useState } from 'react';
import { ResetPasswordSuccess } from './reset-password-success';
import { useQuery } from '@tanstack/react-query';
import { getPingOptions } from '@/client/@tanstack/react-query.gen';
import { client } from '@/client/client.gen';


export function ForgotPasswordForm() {
  const [email, setEmail] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [isSubmitted, setIsSubmitted] = useState(false);
  const router = useRouter();

  const { data, error } = useQuery(getPingOptions());
  client.setConfig({
    baseURL: 'http://localhost:8080',
  });
  console.log("DATA", data);


  const handleReset = async () => {
    if (!email) return;

    setIsLoading(true);
    try {
      // Simulate API call
      await new Promise((resolve) => setTimeout(resolve, 1000));
      setIsSubmitted(true);
    } catch (error) {
      console.error('Reset password error:', error);
      // Handle error (show toast, etc.)
    } finally {
      setIsLoading(false);
    }
  };

  if (isSubmitted) {
    return <ResetPasswordSuccess email={email} onBack={() => router.back()} />;
  }

  return (
    <AuthCard title="Reset Password">
      <IconInput
        label="Email"
        icon={MailIcon}
        placeholder="Enter your email"
        value={email}
        onChangeText={setEmail}
        keyboardType="email-address"
        autoCapitalize="none"
      />

      <Button
        className="mt-2 w-full shadow-md"
        size="lg"
        onPress={handleReset}
        disabled={isLoading || !email}>
        <Text className="font-bold">{isLoading ? 'Sending...' : 'Send Reset Link'}</Text>
      </Button>

      <View className="flex-row items-center justify-center gap-2 pt-4">
        <Text variant="muted">Remember your password?</Text>
        <Link href="/sign-in" asChild>
          <Text variant="small" className="font-semibold text-primary">
            Sign In
          </Text>
        </Link>
      </View>
    </AuthCard>
  );
}

