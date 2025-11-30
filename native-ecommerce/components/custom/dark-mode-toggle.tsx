import { Card, CardContent } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Switch } from '@/components/ui/switch';
import { Text } from '@/components/ui/text';
import { MoonIcon } from 'lucide-react-native';
import { useColorScheme } from 'nativewind';
import { View } from 'react-native';

export function DarkModeToggle() {
  const { colorScheme, toggleColorScheme } = useColorScheme();

  return (
    <Card className="mb-6 overflow-hidden shadow-md shadow-black/5 dark:shadow-black/20">
      <CardContent className="p-4">
        <View className="flex-row items-center">
          <View className="mr-4 rounded-lg bg-muted/30 p-2">
            <Icon as={MoonIcon} size={22} className="text-indigo-500" />
          </View>
          <View className="flex-1">
            <Text className="text-left text-base font-medium">Dark Mode</Text>
            <Text variant="muted" className="text-left text-xs">
              {colorScheme === 'dark' ? 'Currently enabled' : 'Currently disabled'}
            </Text>
          </View>
          <Switch checked={colorScheme === 'dark'} onCheckedChange={toggleColorScheme} />
        </View>
      </CardContent>
    </Card>
  );
}

