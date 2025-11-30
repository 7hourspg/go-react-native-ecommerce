import { ActivityIndicator, View } from 'react-native';
import { Text } from '@/components/ui/text';

interface LoadingProps {
  message?: string;
  size?: 'small' | 'large';
  fullScreen?: boolean;
}

export function Loading({ message, size = 'large', fullScreen = true }: LoadingProps) {
  const content = (
    <View className="items-center justify-center">
      <ActivityIndicator size={size} />
      {message && (
        <Text variant="muted" className="mt-4">
          {message}
        </Text>
      )}
    </View>
  );

  if (fullScreen) {
    return (
      <View className="flex-1 items-center justify-center bg-background">
        {content}
      </View>
    );
  }

  return content;
}

