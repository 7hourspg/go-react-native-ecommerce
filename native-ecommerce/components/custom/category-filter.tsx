import { Button } from '@/components/ui/button';
import { Text } from '@/components/ui/text';
import { ScrollView, View } from 'react-native';

interface CategoryFilterProps {
  categories: string[];
  selectedCategory: string;
  onSelectCategory: (category: string) => void;
}

export function CategoryFilter({ categories, selectedCategory, onSelectCategory }: CategoryFilterProps) {
  return (
    <ScrollView horizontal showsHorizontalScrollIndicator={false}>
      <View className="flex-row gap-2">
        {categories.map((category) => (
          <Button
            key={category}
            variant={selectedCategory === category ? 'default' : 'outline'}
            size="sm"
            className={`rounded-full ${selectedCategory === category ? 'shadow-sm' : ''}`}
            onPress={() => onSelectCategory(category)}>
            <Text className={selectedCategory === category ? 'font-semibold' : ''}>{category}</Text>
          </Button>
        ))}
      </View>
    </ScrollView>
  );
}

