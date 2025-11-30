import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { ShoppingBagIcon } from 'lucide-react-native';
import { View } from 'react-native';

export function EmptyCart() {
  return (
    <View className="flex-1 items-center justify-center p-4">
      <View className="mb-6 h-24 w-24 items-center justify-center rounded-full bg-muted/20">
        <Icon as={ShoppingBagIcon} size={48} className="text-muted-foreground" />
      </View>
      <Text variant="h3" className="mb-2 font-bold">
        Your cart is empty
      </Text>
      <Text variant="muted" className="text-center text-base">
        Add some products to get started
      </Text>
    </View>
  );
}

