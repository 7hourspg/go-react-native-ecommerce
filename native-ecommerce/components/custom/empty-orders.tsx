import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { PackageIcon } from 'lucide-react-native';
import { View } from 'react-native';

export function EmptyOrders() {
  return (
    <View className="flex-1 items-center justify-center p-4">
      <View className="mb-6 h-24 w-24 items-center justify-center rounded-full bg-muted/20">
        <Icon as={PackageIcon} size={48} className="text-muted-foreground" />
      </View>
      <Text variant="h3" className="mb-2 font-bold">
        No orders yet
      </Text>
      <Text variant="muted" className="text-center text-base">
        Your order history will appear here
      </Text>
    </View>
  );
}

