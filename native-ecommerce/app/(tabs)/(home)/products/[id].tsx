import { ProductDetail } from '@/components/products/product-detail';
import { ErrorPage } from '@/components/custom/error-page';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useLocalSearchParams, useRouter } from 'expo-router';

export default function ProductDetailPage() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const router = useRouter();

  if (!id || isNaN(Number(id))) {
    return (
      <SafeAreaView className="flex-1 bg-background" edges={['top']}>
        <ErrorPage
          title="Invalid Product ID"
          message="The product ID you provided is invalid. Please check the URL and try again."
          onGoBack={() => router.back()}
          showRetryButton={false}
        />
      </SafeAreaView>
    );
  }

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <ProductDetail productId={Number(id)} />
    </SafeAreaView>
  );
}

