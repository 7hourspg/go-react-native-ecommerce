import { ProductCard } from '@/components/custom/product-card';
import { SectionHeader } from '@/components/custom/section-header';
import { Text } from '@/components/ui/text';
import { Loading } from '@/components/custom/loading';
import { ErrorPage } from '@/components/custom/error-page';
import { Pressable, ScrollView, View, useWindowDimensions } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import { getProductsOptions, getProductsFeaturedOptions } from '@/client/@tanstack/react-query.gen';
import { useRouter } from 'expo-router';

export default function HomeScreen() {
  const router = useRouter();
  const { width } = useWindowDimensions();

  const {
    data: products,
    isLoading: productsLoading,
    error: productsError,
    refetch: productsRefetch,
  } = useQuery(getProductsOptions());
  const {
    data: featuredProducts,
    isLoading: featuredProductsLoading,
    error: featuredProductsError,
    refetch: featuredProductsRefetch,
  } = useQuery(getProductsFeaturedOptions());

  const padding = 20;
  const gap = 12;
  const cardWidth = (width - padding * 2 - gap) / 2;

  if (productsLoading || featuredProductsLoading) {
    return <Loading message="Loading products..." />;
  }

  if (productsError || featuredProductsError) {
    return (
      <ErrorPage
        title="Error loading products"
        message={productsError?.message || featuredProductsError?.message}
        onRetry={() => {
          productsRefetch();
          featuredProductsRefetch();
        }}
        showBackButton={false}
      />
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      {/* Header */}
      <View className="px-5">
        <View className="mb-3 flex-row items-center justify-between">
          <View>
            <Text variant="h1" className="mb-1 text-left text-3xl font-bold">
              Discover
            </Text>
            <Text variant="muted" className="text-base">
              Find your perfect products
            </Text>
          </View>
        </View>
      </View>

      <ScrollView className="flex-1 px-5 pb-4 pt-4" showsVerticalScrollIndicator={false}>
        {/* Featured Products */}
        <View>
          <SectionHeader
            title="Featured"
            actionText="See all"
            onActionPress={() => router.push('/products')}
          />
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            className="mb-4"
            contentContainerStyle={{ paddingRight: 20 }}>
            <View className="flex-row gap-4">
              {featuredProducts?.data?.map((product) => (
                <Pressable
                  key={product.id}
                  onPress={() => router.push(`/products/${product.id}`)}>
                  <ProductCard {...product} variant="featured" />
                </Pressable>
              ))}
            </View>
          </ScrollView>
        </View>

        {/* All Products Grid */}
        <View>
          <SectionHeader
            title="All Products"
            actionText="View all"
            onActionPress={() => router.push('/products')}
          />
          <View className="flex-row flex-wrap" style={{ gap }}>
            {products?.data?.map((product) => (
              <Pressable
                key={product.id}
                style={{ width: cardWidth }}
                onPress={() => router.push(`/products/${product.id}`)}>
                <ProductCard {...product} variant="grid" />
              </Pressable>
            ))}
          </View>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}
