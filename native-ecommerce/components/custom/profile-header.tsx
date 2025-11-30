import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Card, CardHeader } from '@/components/ui/card';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { UserIcon } from 'lucide-react-native';
import { View } from 'react-native';

interface ProfileHeaderProps {
  name: string;
  email: string;
}

export function ProfileHeader({ name, email }: ProfileHeaderProps) {
  return (
    <Card className="mb-8 overflow-hidden shadow-lg shadow-black/10 dark:shadow-black/30">
      <CardHeader className="pb-6">
        <View className="items-center">
          <View className="mb-4 h-28 w-28 items-center justify-center rounded-full border-4 border-primary/20 bg-primary/10">
            <Avatar className="h-24 w-24">
              <AvatarFallback className="bg-primary/20">
                <Icon as={UserIcon} size={40} className="text-primary" />
              </AvatarFallback>
            </Avatar>
          </View>
          <Text variant="h3" className="mb-2 text-2xl font-bold">
            {name}
          </Text>
          <Text variant="muted" className="text-base">
            {email}
          </Text>
        </View>
      </CardHeader>
    </Card>
  );
}

