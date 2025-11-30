import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { UserIcon } from 'lucide-react-native';
import { useRouter } from 'expo-router';

interface ProfileButtonProps {
  size?: 'sm' | 'md' | 'lg';
}

export function ProfileButton({ size = 'md' }: ProfileButtonProps) {
  const router = useRouter();

  const sizeClasses = {
    sm: 'h-8 w-8',
    md: 'h-10 w-10',
    lg: 'h-12 w-12',
  };

  const iconSizes = {
    sm: 16,
    md: 18,
    lg: 20,
  };

  return (
    <Button
      variant="ghost"
      size="icon"
      className={`${sizeClasses[size]} rounded-full`}
      onPress={() => router.push('/(tabs)/profile')}>
      <Avatar className={`${sizeClasses[size]} border-2 border-primary/20`}>
        <AvatarFallback className="bg-primary/10">
          <Icon as={UserIcon} size={iconSizes[size]} className="text-primary" />
        </AvatarFallback>
      </Avatar>
    </Button>
  );
}

