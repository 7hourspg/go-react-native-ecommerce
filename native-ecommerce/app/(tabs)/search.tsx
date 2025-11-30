import { CategoryFilter } from '@/components/custom/category-filter';
import { ProductCard } from '@/components/custom/product-card';
import { SearchBar } from '@/components/custom/search-bar';
import { Separator } from '@/components/ui/separator';
import { Text } from '@/components/ui/text';
import { useState, useEffect, useMemo } from 'react';
import { ScrollView, View, Pressable } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useQuery } from '@tanstack/react-query';
import {
  getProductsSearchOptions,
  getProductsCategoryByCategoryOptions,
  getProductsOptions,
  getProductsQueryKey,
  getProductsSearchQueryKey,
  getProductsCategoryByCategoryQueryKey,
} from '@/client/@tanstack/react-query.gen';
import { Loading } from '@/components/custom/loading';
import { ErrorPage } from '@/components/custom/error-page';
import { useRouter } from 'expo-router';
import { ModelsProduct } from '@/client/types.gen';

const categories = ['All', 'Electronics', 'Office', 'Home', 'Accessories'];

export default function SearchScreen() {
  const router = useRouter();
  const [searchQuery, setSearchQuery] = useState('');
  const [debouncedQuery, setDebouncedQuery] = useState('');
  const [selectedCategory, setSelectedCategory] = useState('All');

  // Debounce search query
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedQuery(searchQuery);
    }, 500);

    return () => clearTimeout(timer);
  }, [searchQuery]);

  // Get all products as fallback
  const {
    data: allProductsData,
    isLoading: allProductsLoading,
    error: allProductsError,
    refetch: allProductsRefetch,
  } = useQuery({
    ...getProductsOptions(),
    enabled: !debouncedQuery.trim(),
    queryKey: getProductsQueryKey(),
  });

  // Search products when query exists
  const {
    data: searchData,
    isLoading: searchLoading,
    error: searchError,
    refetch: searchRefetch,
  } = useQuery({
    ...getProductsSearchOptions({
      query: {
        query: debouncedQuery,
      },
    }),
    enabled: debouncedQuery.trim().length > 0,
    queryKey: getProductsSearchQueryKey({
      query: {
        query: debouncedQuery,
      },
    }),
  });

  // Get products by category when selected
  const {
    data: categoryData,
    isLoading: categoryLoading,
    error: categoryError,
    refetch: categoryRefetch,
  } = useQuery({
    ...getProductsCategoryByCategoryOptions({
      path: {
        category: selectedCategory,
      },
    }),
    enabled: selectedCategory !== 'All' && !debouncedQuery.trim(),
    queryKey: getProductsCategoryByCategoryQueryKey({
      path: {
        category: selectedCategory,
      },
    }),
  });

  //  which data to display
  const displayData = useMemo(() => {
    if (debouncedQuery.trim()) {
      return searchData?.data || [];
    }
    if (selectedCategory !== 'All') {
      return categoryData?.data || [];
    }
    return allProductsData?.data || [];
  }, [debouncedQuery, selectedCategory, searchData, categoryData, allProductsData]);

  const isLoading = searchLoading || categoryLoading || allProductsLoading;
  const error = searchError || categoryError || allProductsError;

  if (error) {
    return (
      <ErrorPage
        title="Error loading products"
        message={error.message || 'Failed to load products. Please try again.'}
        onRetry={() => {
          searchRefetch();
          categoryRefetch();
          allProductsRefetch();
        }}
        showBackButton={false}
      />
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <View className="flex-1">
        {/* Search Header */}
        <View className="bg-card px-5 pb-4 pt-6 shadow-sm">
          <SearchBar value={searchQuery} onChangeText={setSearchQuery} />
          <CategoryFilter
            categories={categories}
            selectedCategory={selectedCategory}
            onSelectCategory={setSelectedCategory}
          />
        </View>

        <Separator />

        {/* Results */}
        {isLoading ? (
          <Loading message="Searching products..." />
        ) : (
          <ScrollView className="flex-1" showsVerticalScrollIndicator={false}>
            <View className="px-5 pt-6">
              <Text variant="small" className="mb-5 text-muted-foreground">
                {displayData.length} result{displayData.length !== 1 ? 's' : ''} found
                {debouncedQuery.trim() && ` for "${debouncedQuery}"`}
              </Text>

              {displayData.length === 0 ? (
                <View className="py-12">
                  <Text variant="h3" className="mb-2 text-center font-bold">
                    No products found
                  </Text>
                  <Text variant="muted" className="text-center">
                    Try adjusting your search or filters
                  </Text>
                </View>
              ) : (
                <View className="gap-4 pb-6">
                  {displayData.map((product: ModelsProduct) => (
                    <Pressable
                      key={product.id}
                      onPress={() => router.push(`/products/${product.id}`)}>
                      <ProductCard {...product} variant="search" />
                    </Pressable>
                  ))}
                </View>
              )}
            </View>
          </ScrollView>
        )}
      </View>
    </SafeAreaView>
  );
}
