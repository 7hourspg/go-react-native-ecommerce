import { Icon } from '@/components/ui/icon';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { LucideIcon } from 'lucide-react-native';
import { View } from 'react-native';

interface IconInputProps {
  label: string;
  icon: LucideIcon;
  placeholder: string;
  value: string;
  onChangeText: (text: string) => void;
  onBlur?: () => void;
  keyboardType?: 'default' | 'email-address' | 'numeric' | 'phone-pad';
  autoCapitalize?: 'none' | 'sentences' | 'words' | 'characters';
  secureTextEntry?: boolean;
  rightElement?: React.ReactNode;
}

export function IconInput({
  label,
  icon,
  placeholder,
  value,
  onChangeText,
  onBlur,
  keyboardType = 'default',
  autoCapitalize = 'sentences',
  secureTextEntry = false,
  rightElement,
}: IconInputProps) {
  return (
    <View className="gap-2">
      <Label>{label}</Label>
      <View className="relative">
        <View className="absolute left-3 top-3 z-10">
          <Icon as={icon} size={20} className="text-muted-foreground" />
        </View>
        <Input
          placeholder={placeholder}
          value={value}
          onChangeText={onChangeText}
          onBlur={onBlur}
          keyboardType={keyboardType}
          autoCapitalize={autoCapitalize}
          secureTextEntry={secureTextEntry}
          className={rightElement ? 'pl-11 pr-11' : 'pl-11'}
        />
        {rightElement && (
          <View className="absolute right-1 top-1">{rightElement}</View>
        )}
      </View>
    </View>
  );
}

