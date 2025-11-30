import { useQuery } from '@tanstack/react-query';
import { getProductsOptions } from '@/client/@tanstack/react-query.gen';
import { ProductCard } from '@/components/custom/product-card';
import { SectionHeader } from '@/components/custom/section-header';
import { ErrorPage } from '@/components/custom/error-page';
import { Loading } from '@/components/custom/loading';
import { Text } from '@/components/ui/text';
import { View, ScrollView, RefreshControl, Pressable, useWindowDimensions } from 'react-native';
import { useRouter } from 'expo-router';
import { useState } from 'react';

export function ProductsList() {
  const [refreshing, setRefreshing] = useState(false);
  const router = useRouter();
  const { width } = useWindowDimensions();
  const { data, isLoading, error, refetch } = useQuery(getProductsOptions());

  const padding = 20;
  const gap = 16;
  const cardWidth = (width - padding * 2 - gap) / 2;

  const handleRefresh = async () => {
    setRefreshing(true);
    await refetch();
    setRefreshing(false);
  };

  if (isLoading) {
    return <Loading message="Loading products..." />;
  }

  if (error) {
    return (
      <ErrorPage
        title="Error loading products"
        message={
          error.message || 'Failed to load products. Please check your connection and try again.'
        }
        onRetry={() => refetch()}
        showBackButton={false}
      />
    );
  }

  const products = data?.data || [];

  if (products.length === 0) {
    return (
      <View className="flex-1 items-center justify-center bg-background px-5">
        <Text variant="h3" className="mb-2 text-center text-xl font-bold">
          No products found
        </Text>
        <Text variant="muted" className="text-center">
          Check back later for new products
        </Text>
      </View>
    );
  }

  return (
    <View className="flex-1 px-5 pt-6">
      <SectionHeader title="All Products" className="mb-0" />
      <Text variant="small" className="mb-2 text-sm text-muted-foreground">
        {products.length} {products.length === 1 ? 'product' : 'products'} available
      </Text>

      <ScrollView
        showsVerticalScrollIndicator={false}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={handleRefresh} />}>
        <View className="flex-row flex-wrap pb-6" style={{ gap }}>
          {products.map((product) => (
            <Pressable
              key={product.id}
              onPress={() => router.push(`/products/${product.id}`)}
              style={{ width: cardWidth }}>
              <ProductCard {...product} variant="grid" />
            </Pressable>
          ))}
        </View>
      </ScrollView>
    </View>
  );
}
