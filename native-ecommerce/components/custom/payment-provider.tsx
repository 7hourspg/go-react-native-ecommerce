import React from 'react';
import { StripeProvider } from '@stripe/stripe-react-native';
import { Platform } from 'react-native';

export function PaymentProvider({ children }: { children: React.JSX.Element }) {
  const KEY = process.env.EXPO_PUBLIC_STRIPE_PUBLISHABLE_KEY;
  if (!KEY) {
    throw new Error('STRIPE_PUBLISHABLE_KEY is not set');
  }

  // Skip Stripe provider on web
  if (Platform.OS === 'web') {
    return children;
  }

  return <StripeProvider publishableKey={KEY}>{children}</StripeProvider>;
}
