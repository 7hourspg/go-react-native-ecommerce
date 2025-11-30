import { Text } from '@/components/ui/text';
import { View } from 'react-native';

interface AuthHeaderProps {
  title: string;
  subtitle: string;
}

export function AuthHeader({ title, subtitle }: AuthHeaderProps) {
  return (
    <View className="mb-8 items-center">
      <Text variant="h1" className="mb-2 text-4xl font-bold">
        {title}
      </Text>
      <Text variant="muted" className="text-center text-base">
        {subtitle}
      </Text>
    </View>
  );
}

