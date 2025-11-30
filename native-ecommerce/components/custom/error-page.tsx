import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { View } from 'react-native';
import { AlertTriangleIcon, ArrowLeftIcon, RefreshCwIcon } from 'lucide-react-native';
import { useRouter } from 'expo-router';

interface ErrorPageProps {
  title?: string;
  message?: string;
  onRetry?: () => void;
  onGoBack?: () => void;
  showBackButton?: boolean;
  showRetryButton?: boolean;
  retryLabel?: string;
  backLabel?: string;
}

export function ErrorPage({
  title = 'Something went wrong',
  message = 'We encountered an unexpected error. Please try again or go back.',
  onRetry,
  onGoBack,
  showBackButton = true,
  showRetryButton = true,
  retryLabel = 'Try Again',
  backLabel = 'Go Back',
}: ErrorPageProps) {
  const router = useRouter();

  const handleGoBack = () => {
    if (onGoBack) {
      onGoBack();
    } else {
      router.back();
    }
  };

  const handleRetry = () => {
    if (onRetry) {
      onRetry();
    }
  };

  return (
    <View className="flex-1 items-center justify-center bg-background px-5">
      {/* Error Icon */}
      <View className="mb-6 h-24 w-24 items-center justify-center rounded-full bg-red-500/10">
        <Icon as={AlertTriangleIcon} size={48} className="text-red-500" color="#ef4444" />
      </View>

      {/* Title */}
      <Text variant="h1" className="mb-3 text-center text-3xl font-bold">
        {title}
      </Text>

      {/* Message */}
      <Text variant="muted" className="mb-8 text-center text-base leading-6">
        {message}
      </Text>

      {/* Action Buttons */}
      <View className="w-full gap-3">
        {showRetryButton && onRetry && (
          <Button className="w-full shadow-lg" size="lg" onPress={handleRetry}>
            <Icon as={RefreshCwIcon} size={20} />
            <Text className="font-bold">{retryLabel}</Text>
          </Button>
        )}

        {showBackButton && (
          <Button variant="outline" className="w-full border-2" size="lg" onPress={handleGoBack}>
            <Icon as={ArrowLeftIcon} size={20} />
            <Text className="font-semibold">{backLabel}</Text>
          </Button>
        )}
      </View>
    </View>
  );
}

