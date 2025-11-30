import { Text } from '@/components/ui/text';
import { cn } from '@/lib/utils';
import { View } from 'react-native';

interface SectionHeaderProps {
  title: string;
  actionText?: string;
  onActionPress?: () => void;
  className?: string;
}

export function SectionHeader({ title, actionText, onActionPress, className }: SectionHeaderProps) {
  return (
    <View className={cn('mb-5 flex-row items-center justify-between', className)}>
      <Text variant="h3" className="text-xl font-bold">
        {title}
      </Text>
      {actionText && (
        <Text variant="muted" className="text-sm" onPress={onActionPress}>
          {actionText}
        </Text>
      )}
    </View>
  );
}
