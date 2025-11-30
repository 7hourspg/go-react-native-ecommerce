import { useQuery } from '@tanstack/react-query';
import {
  getProductsByIdOptions,
  getWishlistByProductIdOptions,
} from '@/client/@tanstack/react-query.gen';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Separator } from '@/components/ui/separator';
import { Text } from '@/components/ui/text';
import { ErrorPage } from '@/components/custom/error-page';
import { Loading } from '@/components/custom/loading';
import { View, ScrollView, RefreshControl, Alert } from 'react-native';
import { ShoppingCartIcon, HeartIcon, StarIcon, ArrowLeftIcon } from 'lucide-react-native';
import { useMemo, useState } from 'react';
import { useRouter } from 'expo-router';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useAddToCart, useAddToWishlist } from '@/api';

interface ProductDetailProps {
  productId: number;
}

export function ProductDetail({ productId }: ProductDetailProps) {
  const [refreshing, setRefreshing] = useState(false);
  const router = useRouter();

  const { mutateAsync: addToWishlist } = useAddToWishlist();
  const { mutateAsync: addToCart } = useAddToCart();
  const { data, isLoading, error, refetch } = useQuery(
    getProductsByIdOptions({
      path: {
        id: productId,
      },
    })
  );

  const { data: wishlistCheck } = useQuery(
    getWishlistByProductIdOptions({
      path: {
        product_id: productId,
      },
    })
  );

  const isInWishlist = useMemo(() => wishlistCheck?.in_wishlist ?? false, [wishlistCheck]);

  const handleRefresh = async () => {
    setRefreshing(true);
    await refetch();
    setRefreshing(false);
  };

  const handleAddToCart = async () => {
    if (!product) return;

    addToCart({
      body: {
        product_id: productId,
        quantity: 1,
      },
    });
    Alert.alert('Success', 'Product added to cart!');
  };

  const handleWishlist = async () => {
    try {
      await addToWishlist({
        body: {
          product_id: productId,
        },
      });
    } catch (error) {
      Alert.alert('Error', 'Failed to update wishlist');
    }
  };

  if (isLoading) {
    return <Loading message="Loading product..." />;
  }

  if (error) {
    return (
      <ErrorPage
        title="Product not found"
        message={error.message || 'The product you are looking for could not be found.'}
        onRetry={() => refetch()}
        onGoBack={() => router.back()}
      />
    );
  }

  const product = data?.data;

  if (!product) {
    return (
      <ErrorPage
        title="Product not found"
        message="The product you are looking for does not exist or has been removed."
        onGoBack={() => router.back()}
        showRetryButton={false}
      />
    );
  }

  return (
    <View className="flex-1 bg-background">
      {/* Header with Back Button */}
      <View className="ml-3 px-5 pt-4">
        <Button variant="ghost" size="icon" className="h-10 w-10" onPress={() => router.back()}>
          <Icon as={ArrowLeftIcon} size={20} />
          <Text className="font-semibold">Back</Text>
        </Button>
      </View>

      <ScrollView
        className="flex-1"
        showsVerticalScrollIndicator={false}
        refreshControl={<RefreshControl refreshing={refreshing} onRefresh={handleRefresh} />}
        contentContainerStyle={{ paddingBottom: 100 }}>
        <View className="px-5">
          {/* Product Image */}
          <Card className="mb-6 overflow-hidden border-0 shadow-lg shadow-black/10 dark:shadow-black/30">
            <CardContent className="items-center justify-center p-8">
              <View className="mb-4 h-72 w-full items-center justify-center rounded-3xl bg-muted/20">
                <Text className="text-9xl">{product.image || '📦'}</Text>
              </View>
              {product.badge && (
                <Badge
                  variant="secondary"
                  className={`${product.badge_color || 'bg-blue-500'} border-0 text-white`}>
                  {product.badge}
                </Badge>
              )}
            </CardContent>
          </Card>

          {/* Product Info */}
          <View className="mb-6">
            <Text variant="h1" className="mb-3 text-3xl font-bold">
              {product.name}
            </Text>

            <View className="mb-4 flex-row items-center gap-3">
              <View className="flex-row items-center gap-1.5 rounded-full bg-yellow-500/10 px-3 py-1.5">
                <Icon
                  as={StarIcon}
                  size={18}
                  className="text-yellow-500"
                  color="#eab308"
                  fill="#eab308"
                />
                <Text variant="small" className="text-base font-bold">
                  {product.rating?.toFixed(1) || '0.0'}
                </Text>
              </View>
              {product.category && (
                <>
                  <Text variant="muted" className="text-lg">
                    •
                  </Text>
                  <Text variant="muted" className="text-base">
                    {product.category}
                  </Text>
                </>
              )}
            </View>

            <View className="mb-4 flex-row items-baseline gap-3">
              <Text variant="h2" className="text-4xl font-bold text-primary">
                ${product.price || 0}
              </Text>
              {product.original_price && product.original_price > (product.price || 0) && (
                <Text variant="small" className="text-xl text-muted-foreground line-through">
                  ${product.original_price}
                </Text>
              )}
            </View>

            {product.stock !== undefined && (
              <View className="mb-4 flex-row items-center gap-2">
                <View
                  className={`h-2 w-2 rounded-full ${
                    product.stock > 0 ? 'bg-green-500' : 'bg-red-500'
                  }`}
                />
                <Text
                  variant="small"
                  className={`font-semibold ${
                    product.stock > 0
                      ? 'text-green-600 dark:text-green-400'
                      : 'text-red-600 dark:text-red-400'
                  }`}>
                  {product.stock > 0 ? `In Stock (${product.stock} available)` : 'Out of Stock'}
                </Text>
              </View>
            )}
          </View>

          <Separator className="my-6" />

          {/* Description */}
          {product.description && (
            <View className="mb-6">
              <Text variant="h3" className="mb-3 text-xl font-bold">
                Description
              </Text>
              <Text variant="muted" className="text-base leading-7">
                {product.description}
              </Text>
            </View>
          )}
        </View>
      </ScrollView>

      {/* Sticky Action Buttons */}
      <View className="bg-background/95 backdrop-blur-sm">
        <View className="flex-row gap-3 px-5 py-4">
          <Button variant="outline" className="flex-1 shadow-lg" size="lg" onPress={handleWishlist}>
            <Icon
              as={HeartIcon}
              size={20}
              color={isInWishlist ? '#ff0000' : undefined}
              fill={isInWishlist ? '#ff0000' : 'none'}
            />
            <Text className="font-semibold">
              {isInWishlist ? 'In Wishlist' : 'Add to Wishlist'}
            </Text>
          </Button>
          <Button
            className="flex-1 shadow-lg"
            size="lg"
            onPress={handleAddToCart}
            disabled={product.stock === 0}>
            <Icon as={ShoppingCartIcon} size={20} color="white" />
            <Text className="font-bold">Add to Cart</Text>
          </Button>
        </View>
      </View>
    </View>
  );
}
