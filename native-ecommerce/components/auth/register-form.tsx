import { Button } from '@/components/ui/button';
import { IconInput } from './icon-input';
import { PasswordInput } from './password-input';
import { AuthCard } from './auth-card';
import { AuthFooter } from './auth-footer';
import { MailIcon, UserIcon } from 'lucide-react-native';
import { Text } from '@/components/ui/text';
import { useState } from 'react';
import { useRouter } from 'expo-router';
import { useAuthState } from '@/context/auth-context';
import { z } from 'zod';
import { zHandlersRegisterPayload } from '@/client/zod.gen';
import { zodResolver } from '@hookform/resolvers/zod';
import { Controller, useForm } from 'react-hook-form';
import { View } from 'react-native';
import { postAuthRegisterMutation } from '@/client/@tanstack/react-query.gen';
import { useMutation } from '@tanstack/react-query';

const RegisterSchema = zHandlersRegisterPayload
  .extend({
    confirmPassword: z.string().min(8, 'Password must be at least 8 characters'),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: 'Passwords do not match',
    path: ['confirmPassword'],
  });

export function RegisterForm() {
  const { mutateAsync: registerMutation, isPending } = useMutation(postAuthRegisterMutation());
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const { login } = useAuthState();
  const router = useRouter();

  const form = useForm<z.infer<typeof RegisterSchema>>({
    resolver: zodResolver(RegisterSchema),
    defaultValues: {
      name: '',
      email: '',
      password: '',
      confirmPassword: '',
    },
  });

  const onSubmit = async (data: z.infer<typeof RegisterSchema>) => {
    try {
      const response = await registerMutation({
        body: {
          name: data.name,
          email: data.email,
          password: data.password,
        },
      });

      await login(response);

      router.replace('/(tabs)');
    } catch (error) {
      console.error('Registration error:', error);
    }
  };

  return (
    <AuthCard title="Sign Up">
      <Controller
        name="name"
        control={form.control}
        render={({ field: { onChange, onBlur, value }, fieldState: { error } }) => (
          <View>
            <IconInput
              label="Full Name"
              icon={UserIcon}
              placeholder="Enter your full name"
              value={value ?? ''}
              onChangeText={onChange}
              onBlur={onBlur}
              autoCapitalize="words"
            />
            {error && <Text className="mt-1 text-sm text-destructive">{error.message}</Text>}
          </View>
        )}
      />

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
              placeholder="Create a password"
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

      <Controller
        name="confirmPassword"
        control={form.control}
        render={({ field: { onChange, onBlur, value }, fieldState: { error } }) => (
          <View>
            <PasswordInput
              label="Confirm Password"
              placeholder="Confirm your password"
              value={value ?? ''}
              onChangeText={onChange}
              onBlur={onBlur}
              showPassword={showConfirmPassword}
              onToggleShowPassword={() => setShowConfirmPassword(!showConfirmPassword)}
            />
            {error && <Text className="mt-1 text-sm text-destructive">{error.message}</Text>}
          </View>
        )}
      />

      <Button
        className="mt-2 w-full shadow-md"
        size="lg"
        onPress={form.handleSubmit(onSubmit)}
        disabled={isPending}>
        <Text className="font-bold">{isPending ? 'Creating Account...' : 'Sign Up'}</Text>
      </Button>

      <AuthFooter question="Already have an account?" linkText="Sign In" linkHref="/sign-in" />
    </AuthCard>
  );
}
