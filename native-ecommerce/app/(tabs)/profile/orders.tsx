import { OrderItemCard } from '@/components/custom/order-item-card';
import { EmptyOrders } from '@/components/custom/empty-orders';
import { Text } from '@/components/ui/text';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { getOrdersOptions, getOrdersQueryKey } from '@/client/@tanstack/react-query.gen';
import { ErrorPage } from '@/components/custom/error-page';
import { Loading } from '@/components/custom/loading';

export default function OrdersScreen() {
  const { data: ordersData, isLoading, error, refetch } = useQuery({
    ...getOrdersOptions(),
    queryKey: getOrdersQueryKey(),
  });

  if (isLoading) {
    return <Loading message="Loading orders..." />;
  }
  if (error) {
    return (
      <ErrorPage
        title="Error loading orders"
        message={
          error.message ||
          'Failed to load orders. Please check your connection and try again.'
        }
        onRetry={() => refetch()}
        showBackButton={false}
      />
    );
  }

  const orders = Array.isArray(ordersData?.data) ? ordersData.data : [];

  if (orders.length === 0) {
    return <EmptyOrders />;
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <View className="flex-1">
        <View className="flex-1 px-5 pt-6">
          <Text variant="h2" className="mb-2 border-none text-2xl font-bold">
            My Orders
          </Text>
          <ScrollView showsVerticalScrollIndicator={false}>
            <View className="gap-4 pb-6">
              {orders.map((order) => {
                return <OrderItemCard key={order.id} {...order} />;
              })}
            </View>
          </ScrollView>
        </View>
      </View>
    </SafeAreaView>
  );
}

