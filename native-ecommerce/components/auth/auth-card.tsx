import { Card, CardContent, CardHeader } from '@/components/ui/card';
import { Text } from '@/components/ui/text';
import { ReactNode } from 'react';

interface AuthCardProps {
  title: string;
  children: ReactNode;
}

export function AuthCard({ title, children }: AuthCardProps) {
  return (
    <Card className="overflow-hidden shadow-lg shadow-black/10 dark:shadow-black/30">
      <CardHeader className="pb-6">
        <Text variant="h3" className="text-center text-2xl font-bold">
          {title}
        </Text>
      </CardHeader>
      <CardContent className="gap-4">{children}</CardContent>
    </Card>
  );
}

