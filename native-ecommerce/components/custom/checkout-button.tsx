import { useStripe } from '@stripe/stripe-react-native';
import { useState } from 'react';
import { Alert } from 'react-native';
import { Button } from '@/components/ui/button';
import { Text } from '@/components/ui/text';
import { router } from 'expo-router';
import { postPaymentsConfirmSuccessByOrderId } from '@/client/sdk.gen';
import { useMutation } from '@tanstack/react-query';
import { postOrdersCheckoutMutation, getCartQueryKey } from '@/client/@tanstack/react-query.gen';

export default function CheckoutButton() {
  const { initPaymentSheet, presentPaymentSheet } = useStripe();
  const [loading, setLoading] = useState(false);

  const { mutateAsync: checkout } = useMutation({
    ...postOrdersCheckoutMutation(),
    async onSuccess(data, _, __, context) {
      const { client_secret, order } = data;

      if (!client_secret) {
        Alert.alert('Error', 'No payment secret received');
        setLoading(false);
        return;
      }

      // Initialize payment sheet
      const { error: initError } = await initPaymentSheet({
        merchantDisplayName: 'E-Commerce Store',
        paymentIntentClientSecret: client_secret,
        defaultBillingDetails: {
          name: 'Customer',
        },
      });

      if (initError) {
        Alert.alert('Error', initError.message);
        setLoading(false);
        return;
      }

      // Present payment sheet
      const { error: paymentError } = await presentPaymentSheet();

      if (paymentError) {
        Alert.alert('Payment Cancelled', paymentError.message);
        setLoading(false);
      } else {
        // Confirm payment success on backend
        try {
          await postPaymentsConfirmSuccessByOrderId({
            path: {
              orderId: order?.id?.toString() ?? '',
            },
          });

          context?.client.invalidateQueries({ queryKey: getCartQueryKey() });

          Alert.alert('Success!', 'Payment confirmed successfully!', [
            {
              text: 'OK',
              onPress: () => router.push('/(tabs)/(home)'),
            },
          ]);
        } catch (error: any) {
          Alert.alert('Error', error.message || 'Failed to confirm payment');
        }
        setLoading(false);
      }
    },
    onError(error) {
      Alert.alert('Error', error.message || 'Something went wrong');
      setLoading(false);
    },
  });

  const handleCheckout = () => {
    setLoading(true);
    checkout({});
  };

  return (
    <Button className="w-full shadow-md" size="lg" onPress={handleCheckout} disabled={loading}>
      <Text className="font-bold">{loading ? 'Processing...' : 'Checkout'}</Text>
    </Button>
  );
}
