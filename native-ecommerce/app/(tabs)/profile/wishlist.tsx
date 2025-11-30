import { WishlistItemCard } from '@/components/custom/wishlist-item-card';
import { EmptyWishlist } from '@/components/custom/empty-wishlist';
import { Text } from '@/components/ui/text';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { getWishlistOptions } from '@/client/@tanstack/react-query.gen';
import { ErrorPage } from '@/components/custom/error-page';
import { Loading } from '@/components/custom/loading';

export default function WishlistScreen() {
  const { data: wishlistData, isLoading, error, refetch } = useQuery(getWishlistOptions());

  if (isLoading) {
    return <Loading message="Loading wishlist..." />;
  }
  if (error) {
    return (
      <ErrorPage
        title="Error loading wishlist"
        message={
          error.message ||
          'Failed to load wishlist items. Please check your connection and try again.'
        }
        onRetry={() => refetch()}
        showBackButton={false}
      />
    );
  }

  const wishlistItems = Array.isArray(wishlistData?.data) ? wishlistData.data : [];

  if (wishlistItems.length === 0) {
    return <EmptyWishlist />;
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <View className="flex-1">
        <View className="flex-1 px-5 pt-6">
          <Text variant="h2" className="mb-2 border-none text-2xl font-bold">
            My Wishlist
          </Text>
          <ScrollView showsVerticalScrollIndicator={false}>
            <View className="gap-4 pb-6">
              {wishlistItems?.map((item) => {
                return <WishlistItemCard key={item.id} {...item} />;
              })}
            </View>
          </ScrollView>
        </View>
      </View>
    </SafeAreaView>
  );
}
