import { ModelsOrder } from '@/client/types.gen';
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { PackageIcon } from 'lucide-react-native';
import { View } from 'react-native';

export function OrderItemCard(data: ModelsOrder) {
  const formatDate = (dateString?: string) => {
    if (!dateString) return 'N/A';
    try {
      const date = new Date(dateString);
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
      const month = months[date.getMonth()];
      const day = date.getDate();
      const year = date.getFullYear();
      return `${month} ${day}, ${year}`;
    } catch {
      return dateString;
    }
  };

  const getStatusColor = (status?: string) => {
    switch (status?.toLowerCase()) {
      case 'delivered':
        return 'bg-green-500';
      case 'shipped':
        return 'bg-blue-500';
      case 'processing':
        return 'bg-yellow-500';
      case 'cancelled':
        return 'bg-red-500';
      default:
        return 'bg-gray-500';
    }
  };

  const getStatusText = (status?: string) => {
    if (!status) return 'Pending';
    return status.charAt(0).toUpperCase() + status.slice(1);
  };

  const itemCount = data.items?.length ?? 0;
  const totalItems = data.items?.reduce((sum, item) => sum + (item.quantity ?? 0), 0) ?? 0;

  return (
    <Card className="overflow-hidden shadow-md shadow-black/5 dark:shadow-black/20">
      <CardHeader className="pb-3">
        <View className="flex-row items-center justify-between">
          <View className="flex-1">
            <View className="mb-2 flex-row items-center gap-2">
              <Icon as={PackageIcon} size={18} className="text-muted-foreground" />
              <Text variant="small" className="text-muted-foreground">
                Order #{data.id}
              </Text>
            </View>
            <Text variant="small" className="text-muted-foreground">
              {formatDate(data.created_at)}
            </Text>
          </View>
          <Badge className={getStatusColor(data.status)}>
            <Text className="text-xs font-semibold text-white">
              {getStatusText(data.status)}
            </Text>
          </Badge>
        </View>
      </CardHeader>
      <CardContent className="pb-3">
        <View className="gap-2">
          <Text variant="small" className="text-muted-foreground">
            {itemCount} item{itemCount !== 1 ? 's' : ''} • {totalItems} total quantity
          </Text>
          {data.items && data.items.length > 0 && (
            <View className="gap-1">
              {data.items.slice(0, 2).map((item) => (
                <View key={item.id} className="flex-row items-center gap-2">
                  <View className="h-10 w-10 items-center justify-center rounded-lg bg-muted/20">
                    <Text className="text-2xl">{item.product?.image}</Text>
                  </View>
                  <View className="flex-1">
                    <Text variant="small" className="mb-1 font-medium" numberOfLines={1}>
                      {item.product?.name}
                    </Text>
                    <Text variant="small" className="text-muted-foreground">
                      Qty: {item.quantity} × ${item.price?.toFixed(2)}
                    </Text>
                  </View>
                </View>
              ))}
              {data.items.length > 2 && (
                <Text variant="small" className="text-muted-foreground">
                  +{data.items.length - 2} more item{data.items.length - 2 !== 1 ? 's' : ''}
                </Text>
              )}
            </View>
          )}
        </View>
      </CardContent>
      <CardFooter className="flex-row items-center justify-between border-t border-border pt-3">
        <View>
          <Text variant="small" className="text-muted-foreground">
            Total
          </Text>
          <Text variant="h4" className="text-xl font-bold text-primary">
            ${data.total?.toFixed(2)}
          </Text>
        </View>
        {data.payment_status && (
          <View className="items-end gap-1">
            <Text variant="small" className="text-muted-foreground">
              Payment
            </Text>
            <Badge
              className={
                data.payment_status === 'succeeded'
                  ? 'bg-green-500'
                  : data.payment_status === 'failed'
                    ? 'bg-red-500'
                    : 'bg-yellow-500'
              }>
              <Text className="text-xs font-semibold text-white">
                {data.payment_status.charAt(0).toUpperCase() + data.payment_status.slice(1)}
              </Text>
            </Badge>
          </View>
        )}
      </CardFooter>
    </Card>
  );
}

