import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { LucideIcon, SettingsIcon } from 'lucide-react-native';
import { View } from 'react-native';

interface ProfileMenuItemProps {
  icon: LucideIcon;
  label: string;
  count?: number;
  color?: string;
  onPress?: () => void;
}

export function ProfileMenuItem({ icon, label, count, color = 'text-blue-500', onPress }: ProfileMenuItemProps) {
  return (
    <Card className="overflow-hidden shadow-md shadow-black/5 dark:shadow-black/20">
      <CardContent className="p-4">
        <Button variant="ghost" className="w-full justify-start" onPress={onPress}>
          <View className={`mr-4 rounded-lg bg-muted/30 p-2`}>
            <Icon as={icon} size={22} className={color} />
          </View>
          <Text className="flex-1 text-left text-base font-medium">{label}</Text>
          {count !== undefined && (
            <Badge variant="default" className="ml-2 min-w-[28px] rounded-full bg-black dark:bg-white px-3 py-1">
              <Text className="text-xs font-semibold text-white dark:text-black">{count}</Text>
            </Badge>
          )}
          <Icon as={SettingsIcon} size={18} className="ml-2 text-muted-foreground" />
        </Button>
      </CardContent>
    </Card>
  );
}

