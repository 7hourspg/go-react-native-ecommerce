import { ModelsWishlist } from '@/client/types.gen';
import { Button } from '@/components/ui/button';
import { Card, CardFooter, CardHeader } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { ShoppingCartIcon, TrashIcon } from 'lucide-react-native';
import { View } from 'react-native';
import { useAddToCart, useRemoveFromWishlist } from '@/api';
import { Alert } from 'react-native';

export function WishlistItemCard(data: ModelsWishlist) {
  const { mutateAsync: addToCart } = useAddToCart();

  const { mutateAsync: removeFromWishlist } = useRemoveFromWishlist();

  const onRemove = async () => {
    await removeFromWishlist({
      path: { product_id: data.product_id ?? 0 },
    });
  };

  const handleAddToCart = async () => {
    if (!data.product_id) return;
    await addToCart({
      body: {
        product_id: data.product_id,
        quantity: 1,
      },
    });
    Alert.alert('Success', 'Product added to cart!');
  };
  return (
    <Card className="overflow-hidden shadow-md shadow-black/5 dark:shadow-black/20">
      <CardHeader className="pb-3">
        <View className="flex-row items-center gap-4">
          <View className="h-20 w-20 items-center justify-center rounded-xl bg-muted/20">
            <Text className="text-5xl">{data.product?.image}</Text>
          </View>
          <View className="flex-1">
            <Text variant="h4" className="mb-2 text-lg font-semibold">
              {data.product?.name}
            </Text>
            <View className="flex-row items-center gap-2">
              <Text variant="h4" className="text-xl font-bold text-primary">
                ${data.product?.price}
              </Text>
              {data.product?.original_price &&
                data.product?.original_price > (data.product?.price ?? 0) && (
                  <Text variant="muted" className="text-sm line-through">
                    ${data.product?.original_price}
                  </Text>
                )}
            </View>
          </View>
          <Button
            variant="ghost"
            size="icon"
            className="h-9 w-9 rounded-full"
            onPress={() => onRemove()}>
            <Icon as={TrashIcon} size={18} className="text-destructive" />
          </Button>
        </View>
      </CardHeader>
      <CardFooter className="px-4 pb-4">
        <Button variant="outline" className="flex-1" size="lg" onPress={handleAddToCart}>
          <Icon as={ShoppingCartIcon} size={18} />
          <Text className="font-semibold">Add to Cart</Text>
        </Button>
      </CardFooter>
    </Card>
  );
}
