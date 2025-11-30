import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import { Text } from '@/components/ui/text';
import { Badge } from '@/components/ui/badge';
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import { ActivityIndicator, View } from 'react-native';
import { ModelsCartSummary } from '@/client/types.gen';
import CheckoutButton from './checkout-button';

export function CheckoutSummary({
  cartItemsData,
}: {
  cartItemsData: ModelsCartSummary & { isLoading: boolean };
}) {
  const { subtotal, shipping, total, tax, isLoading } = cartItemsData;

  return (
    <View className="border-t border-border bg-card px-5 shadow-lg shadow-black/10 py-4">
      <Accordion type="single" collapsible defaultValue="item-1">
        <AccordionItem value="item-1">
          <AccordionTrigger>
            <View className="flex-row items-center gap-2">
              <Text>Order Summary</Text>
              <Badge variant="outline" className="rounded-full bg-green-500 text-xs text-white">
                {shipping === 0 ? 'Free' : `$${shipping?.toFixed(2)}`}
              </Badge>
            </View>
          </AccordionTrigger>
          <AccordionContent>
            {isLoading ? (
              <View className="flex-1 items-center justify-center py-5">
                <ActivityIndicator size="large" />
              </View>
            ) : (
              <View className="mb-5 gap-3">
                <View className="flex-row justify-between">
                  <Text variant="muted" className="text-base">
                    Subtotal
                  </Text>
                  <Text className="text-base font-semibold">${subtotal?.toFixed(2)}</Text>
                </View>
                <View className="flex-row justify-between">
                  <Text variant="muted" className="text-base">
                    Shipping
                  </Text>
                  <Badge variant="outline" className="rounded-full bg-green-500 text-xs text-white">
                    {shipping === 0 ? 'Free' : `$${shipping?.toFixed(2)}`}
                  </Badge>
                </View>
                <View className="flex-row justify-between">
                  <Text variant="muted" className="text-base">
                    Tax
                  </Text>
                  <Text className="text-base font-semibold">${tax?.toFixed(2)}</Text>
                </View>
                <Separator className="my-2" />
                <View className="flex-row justify-between">
                  <Text variant="h4" className="text-lg font-bold">
                    Total
                  </Text>
                  <Text variant="h4" className="text-2xl font-bold text-primary">
                    ${total?.toFixed(2)}
                  </Text>
                </View>
              </View>
            )}
          </AccordionContent>
        </AccordionItem>
      </Accordion>
      <CheckoutButton />
    </View>
  );
}
