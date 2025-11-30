import { Separator } from '@/components/ui/separator';
import { Text } from '@/components/ui/text';
import { Link } from 'expo-router';
import { View } from 'react-native';

interface AuthFooterProps {
  question: string;
  linkText: string;
  linkHref: string;
}

export function AuthFooter({ question, linkText, linkHref }: AuthFooterProps) {
  return (
    <>
      <Separator className="my-4" />
      <View className="flex-row items-center justify-center gap-2">
        <Text variant="muted">{question}</Text>
        <Link href={linkHref} asChild>
          <Text variant="small" className="font-semibold text-primary">
            {linkText}
          </Text>
        </Link>
      </View>
    </>
  );
}

