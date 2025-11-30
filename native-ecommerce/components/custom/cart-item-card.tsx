import { useRemoveFromCart, useUpdateCartItemQuantity } from '@/api';
import { ModelsCartItem } from '@/client/types.gen';
import { Button } from '@/components/ui/button';
import { Card, CardFooter, CardHeader } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { MinusIcon, PlusIcon, TrashIcon } from 'lucide-react-native';
import { View } from 'react-native';

export function CartItemCard(data: ModelsCartItem) {
  const { mutateAsync: updateCartItemQuantity } = useUpdateCartItemQuantity();
  const { mutateAsync: removeFromCart } = useRemoveFromCart();

  const onRemove = async () => {
    await removeFromCart({
      path: { id: data.id ?? 0 },
    });
  };

  const onUpdateQuantity = (id: number, delta: number) => {
    updateCartItemQuantity({
      path: { id: data.id ?? 0 },
      body: { quantity: (data.quantity ?? 0) + delta },
    });
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
            <Text variant="h4" className="text-xl font-bold text-primary">
              ${data.product?.price}
            </Text>
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
        <View className="flex-row items-center gap-4">
          <View className="flex-row items-center gap-1 rounded-lg border border-input bg-muted/20">
            <Button
              variant="ghost"
              size="icon"
              className="h-9 w-9"
              onPress={() => onUpdateQuantity(data.product?.id ?? 0, -1)}>
              <Icon as={MinusIcon} size={16} />
            </Button>
            <Text variant="small" className="min-w-[40px] text-center font-semibold">
              {data.quantity}
            </Text>
            <Button
              variant="ghost"
              size="icon"
              className="h-9 w-9"
              onPress={() => onUpdateQuantity(data.product?.id ?? 0, 1)}>
              <Icon as={PlusIcon} size={16} />
            </Button>
          </View>
          <Text variant="h4" className="flex-1 text-right text-lg font-bold">
            ${(data.product?.price ?? 0 * (data.quantity ?? 0)).toFixed(2)}
          </Text>
        </View>
      </CardFooter>
    </Card>
  );
}
