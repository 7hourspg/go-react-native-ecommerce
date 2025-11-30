import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { Input } from '@/components/ui/input';
import { SearchIcon, XIcon } from 'lucide-react-native';
import { View } from 'react-native';

interface SearchBarProps {
  value: string;
  onChangeText: (text: string) => void;
}

export function SearchBar({ value, onChangeText }: SearchBarProps) {
  return (
    <View className="mb-4 flex-row items-center gap-3">
      <View className="flex-1 flex-row items-center rounded-xl border border-input bg-background px-4 shadow-sm">
        <Icon as={SearchIcon} size={20} className="mr-2 text-muted-foreground" />
        <Input
          placeholder="Search products..."
          value={value}
          onChangeText={onChangeText}
          className="flex-1 border-0 bg-transparent text-base"
        />
        {value.length > 0 && (
          <Button variant="ghost" size="icon" className="h-7 w-7" onPress={() => onChangeText('')}>
            <Icon as={XIcon} size={16} />
          </Button>
        )}
      </View>
      {/* <Button variant="outline" size="icon" className="h-12 w-12 rounded-xl" onPress={onFilterPress}>
        <Icon as={FilterIcon} size={20} />
      </Button> */}
    </View>
  );
}
