import { CartItemCard } from '@/components/custom/cart-item-card';
import { CheckoutSummary } from '@/components/custom/checkout-summary';
import { EmptyCart } from '@/components/custom/empty-cart';
import { Text } from '@/components/ui/text';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { getCartOptions, getCartQueryKey } from '@/client/@tanstack/react-query.gen';
import { ErrorPage } from '@/components/custom/error-page';
import { Loading } from '@/components/custom/loading';

export default function CartScreen() {
  const {
    data: cartItemsData,
    isLoading,
    error,
    refetch,
  } = useQuery({
    ...getCartOptions(),
    queryKey: getCartQueryKey(),
  });

  if (isLoading) {
    return <Loading message="Loading cart items..." />;
  }
  if (error) {
    return (
      <ErrorPage
        title="Error loading cart items"
        message={
          error.message || 'Failed to load cart items. Please check your connection and try again.'
        }
        onRetry={() => refetch()}
        showBackButton={false}
      />
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <View className="flex-1">
        {cartItemsData?.data?.items?.length === 0 ? (
          <EmptyCart />
        ) : (
          <>
            <View className="flex-1 px-5 pt-6">
              <Text variant="h2" className="mb-2 border-none text-2xl font-bold">
                Shopping Cart
              </Text>
              <ScrollView showsVerticalScrollIndicator={false}>
                <View className="gap-4 pb-6">
                  {cartItemsData?.data?.items?.map((item) => {
                    return <CartItemCard key={item.id} {...item} />;
                  })}
                </View>
              </ScrollView>
            </View>

            {/* @ts-ignore */}
            {cartItemsData?.data && <CheckoutSummary cartItemsData={{ ...cartItemsData?.data }} />}
          </>
        )}
      </View>
    </SafeAreaView>
  );
}
