import { DarkModeToggle } from '@/components/custom/dark-mode-toggle';
import { ProfileHeader } from '@/components/custom/profile-header';
import { ProfileMenuItem } from '@/components/custom/profile-menu-item';
import { Button } from '@/components/ui/button';
import { Icon } from '@/components/ui/icon';
import { Text } from '@/components/ui/text';
import { HeartIcon, LogOutIcon, PackageIcon } from 'lucide-react-native';
import { useRouter } from 'expo-router';
import { ScrollView, View } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useAuthState } from '@/context/auth-context';
import { useQuery } from '@tanstack/react-query';
import { getWishlistOptions, getWishlistQueryKey, getOrdersOptions, getOrdersQueryKey } from '@/client/@tanstack/react-query.gen';

export default function ProfileScreen() {
  const router = useRouter();
  const { user, logout } = useAuthState();
  
  const { data: wishlistData } = useQuery({
    ...getWishlistOptions(),
    queryKey: getWishlistQueryKey(),
  });

  const { data: ordersData } = useQuery({
    ...getOrdersOptions(),
    queryKey: getOrdersQueryKey(),
  });

  const menuItems = [
    { 
      icon: PackageIcon, 
      label: 'My Orders', 
      count: ordersData?.data?.length ?? 0, 
      color: 'text-blue-500', 
      onPress: () => router.push('/profile/orders' as any)
    },
    { 
      icon: HeartIcon, 
      label: 'Wishlist', 
      count: wishlistData?.data?.length ?? 0, 
      color: 'text-red-500',
      onPress: () => router.push('/profile/wishlist' as any)
    },
  ];

  const handleLogout = async () => {
    try {
      await logout();
      router.replace('/sign-in');
    } catch (error) {
      console.error('Logout error:', error);
    }
  };

  return (
    <SafeAreaView className="flex-1 bg-background" edges={['top']}>
      <View className="px-5 pt-6">
        <Text variant="h2" className="mb-2 border-none text-2xl font-bold">
          Profile
        </Text>
      </View>
      <ScrollView className="flex-1" showsVerticalScrollIndicator={false}>
        <View className="px-5 pb-4 pt-6">
          <ProfileHeader name={user?.name || 'User'} email={user?.email || 'user@example.com'} />

          {/* Menu Items */}
          <View className="mb-6 gap-3">
            {menuItems.map((item, index) => (
              <ProfileMenuItem key={index} {...item} />
            ))}
          </View>

          <DarkModeToggle />
        </View>
      </ScrollView>

      {/* Logout */}
      <View className="px-5 pb-4 pt-6">
        <Button variant="destructive" className="w-full shadow-md" size="lg" onPress={handleLogout}>
          <Icon as={LogOutIcon} size={20} color="white" />
          <Text className="font-bold">Log Out</Text>
        </Button>
      </View>
    </SafeAreaView>
  );
}
