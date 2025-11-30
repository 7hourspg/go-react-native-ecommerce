import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { HeartIcon, ShoppingCartIcon, StarIcon } from 'lucide-react-native';
import { Alert, View } from 'react-native';
import { ModelsProduct } from '@/client/types.gen';
import { useAddToCart } from '@/api';

interface ProductCardProps extends ModelsProduct {
  variant?: 'featured' | 'grid' | 'search';
  onWishlist?: () => void;
  originalPrice?: number;
}

export function ProductCard({
  name,
  price,
  rating,
  image,
  originalPrice,
  badge,
  id,
  onWishlist,
  variant = 'grid',
}: ProductCardProps) {
  const { mutateAsync: addToCart } = useAddToCart();

  const handleAddToCart = async () => {
    if (!id) return;
    await addToCart({
      body: {
        product_id: id,
        quantity: 1,
      },
    });
    Alert.alert('Success', 'Product added to cart!');
  };

  if (variant === 'featured') {
    return (
      <Card className="w-72 overflow-hidden shadow-lg shadow-black/10 dark:shadow-black/30">
        <CardHeader className="pb-3">
          <View className="flex-row items-start justify-between">
            {badge && (
              <Badge variant="secondary" className="border-0 bg-blue-500 !text-white">
                {badge}
              </Badge>
            )}
            <Button
              variant="ghost"
              size="icon"
              className="h-9 w-9 rounded-full bg-background/50"
              onPress={onWishlist}>
              <Icon as={HeartIcon} size={18} className="text-muted-foreground" />
            </Button>
          </View>
        </CardHeader>
        <CardContent className="items-center pb-4">
          <View className="mb-4 h-32 w-32 items-center justify-center rounded-2xl bg-muted/30">
            <Text className="text-7xl">{image}</Text>
          </View>
          <Text variant="h4" className="mb-2 text-center text-lg font-semibold">
            {name}
          </Text>
          <View className="mb-3 flex-row items-center gap-1">
            <Icon as={StarIcon} size={16} className="text-yellow-500" />
            <Text variant="small" className="font-medium">
              {rating}
            </Text>
          </View>
          <View className="flex-row items-baseline gap-2">
            <Text variant="h3" className="text-2xl font-bold text-primary">
              ${price}
            </Text>
            {originalPrice && (
              <Text variant="small" className="text-muted-foreground line-through">
                ${originalPrice}
              </Text>
            )}
          </View>
        </CardContent>
        <CardFooter className="px-4 pb-4">
          <Button className="flex-1 shadow-md" onPress={handleAddToCart}>
            <Icon as={ShoppingCartIcon} size={18} color="white" />
            <Text className="font-semibold">Add to Cart</Text>
          </Button>
        </CardFooter>
      </Card>
    );
  }

  if (variant === 'search') {
    return (
      <Card className="overflow-hidden shadow-md shadow-black/5 dark:shadow-black/20">
        <CardHeader className="pb-3">
          <View className="flex-row items-start gap-4">
            <View className="h-20 w-20 items-center justify-center rounded-xl bg-muted/20">
              <Text className="text-4xl">{image}</Text>
            </View>
            <View className="flex-1">
              <Text variant="h4" className="mb-2 text-lg font-semibold">
                {name}
              </Text>
              <View className="mb-3 flex-row items-center gap-2">
                <View className="flex-row items-center gap-1">
                  <Icon as={StarIcon} size={14} className="text-yellow-500" />
                  <Text variant="small" className="font-medium">
                    {rating}
                  </Text>
                </View>
              </View>
              <Text variant="h4" className="text-xl font-bold text-primary">
                ${price}
              </Text>
            </View>
          </View>
        </CardHeader>
        <CardFooter className="px-4 pb-4">
          <Button variant="outline" className="flex-1" onPress={handleAddToCart}>
            <Icon as={ShoppingCartIcon} size={18} />
            <Text className="font-semibold">Add to Cart</Text>
          </Button>
        </CardFooter>
      </Card>
    );
  }

  // Grid variant (default)
  return (
    <Card className="w-full overflow-hidden border border-border/50 shadow-sm shadow-black/5 dark:shadow-black/20">
      <CardHeader className="pb-2">
        <View className="items-center">
          <View className="mb-3 h-28 w-full items-center justify-center rounded-lg bg-muted/30">
            <Text className="text-6xl">{image}</Text>
          </View>
          <Text
            variant="small"
            className="mb-1.5 text-center text-sm font-semibold"
            numberOfLines={2}>
            {name}
          </Text>
          <View className="mb-2 flex-row items-center gap-1">
            <Icon
              as={StarIcon}
              size={14}
              className="text-yellow-500"
              color="#eab308"
              fill="#eab308"
            />
            <Text variant="small" className="text-xs font-semibold">
              {rating?.toFixed(1)}
            </Text>
          </View>
          <View className="w-full flex-row items-baseline justify-center gap-1.5">
            <Text variant="h4" className="text-base font-bold text-foreground">
              ${price}
            </Text>
            {originalPrice && (
              <Text variant="small" className="text-xs text-muted-foreground line-through">
                ${originalPrice}
              </Text>
            )}
          </View>
        </View>
      </CardHeader>
      <CardFooter className="px-3 pb-3 pt-0">
        <Button
          variant="outline"
          size="sm"
          className="flex-1 border-primary/20"
          onPress={handleAddToCart}>
          <Icon as={ShoppingCartIcon} size={16} className="text-primary" />
        </Button>
      </CardFooter>
    </Card>
  );
}
