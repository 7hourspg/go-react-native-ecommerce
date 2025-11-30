import { Button } from '@/components/ui/button';
import { IconInput } from './icon-input';
import { PasswordInput } from './password-input';
import { AuthCard } from './auth-card';
import { AuthFooter } from './auth-footer';
import { MailIcon } from 'lucide-react-native';
import { Link, useRouter } from 'expo-router';
import { Text } from '@/components/ui/text';
import { View } from 'react-native';
import { useState } from 'react';
import { useAuthState } from '@/context/auth-context';
import { zHandlersLoginPayload } from '@/client/zod.gen';
import { zodResolver } from '@hookform/resolvers/zod';
import { Controller, useForm } from 'react-hook-form';
import { HandlersLoginPayload } from '@/client/types.gen';
import { postAuthLoginMutation } from '@/client/@tanstack/react-query.gen';
import { useMutation } from '@tanstack/react-query';

export function SignInForm() {
  const [showPassword, setShowPassword] = useState(false);

  const { mutateAsync: loginMutation, isPending } = useMutation(postAuthLoginMutation());

  const { login } = useAuthState();
  const router = useRouter();

  const form = useForm<HandlersLoginPayload>({
    resolver: zodResolver(zHandlersLoginPayload),
    defaultValues: {
      email: 'rajiv@gmail.com',
      password: '12345678',
    },
  });

  const onSubmit = async (data: HandlersLoginPayload) => {
    try {
      const response = await loginMutation({
        body: {
          email: data.email,
          password: data.password,
        },
      });

      await login(response);
      router.replace('/(tabs)');
    } catch (error) {
      console.error('Sign in error:', error);
    }
  };

  return (
    <AuthCard title="Sign In">
      <Controller
        name="email"
        control={form.control}
        render={({ field: { onChange, onBlur, value }, fieldState: { error } }) => (
          <View>
            <IconInput
              label="Email"
              icon={MailIcon}
              placeholder="Enter your email"
              value={value ?? ''}
              onChangeText={onChange}
              onBlur={onBlur}
              keyboardType="email-address"
              autoCapitalize="none"
            />
            {error && <Text className="mt-1 text-sm text-destructive">{error.message}</Text>}
          </View>
        )}
      />

      <Controller
        name="password"
        control={form.control}
        render={({ field: { onChange, onBlur, value }, fieldState: { error } }) => (
          <View>
            <PasswordInput
              label="Password"
              placeholder="Enter your password"
              value={value ?? ''}
              onChangeText={onChange}
              onBlur={onBlur}
              showPassword={showPassword}
              onToggleShowPassword={() => setShowPassword(!showPassword)}
            />
            {error && <Text className="mt-1 text-sm text-destructive">{error.message}</Text>}
          </View>
        )}
      />

      <View className="flex-row items-center justify-between">
        <View className="flex-row items-center gap-2">
          {/* Checkbox would go here if needed */}
        </View>
        <Link href="/forgot-password" asChild>
          <Text variant="small" className="text-primary">
            Forgot Password?
          </Text>
        </Link>
      </View>

      <Button
        className="mt-2 w-full shadow-md"
        size="lg"
        onPress={form.handleSubmit(onSubmit)}
        disabled={isPending}>
        <Text className="font-bold">{isPending ? 'Signing in...' : 'Sign In'}</Text>
      </Button>

      <AuthFooter question="Don't have an account?" linkText="Sign Up" linkHref="/register" />
    </AuthCard>
  );
}
