import { ProductsList } from '@/components/products/products-list';
import { SafeAreaView } from 'react-native-safe-area-context';

export default function ProductsPage() {
  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <ProductsList />
    </SafeAreaView>
  );
}

