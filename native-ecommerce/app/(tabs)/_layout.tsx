import { Tabs } from 'expo-router';
import { HomeIcon, SearchIcon, ShoppingCartIcon, UserIcon } from 'lucide-react-native';
import { useColorScheme } from 'nativewind';
import { useQuery } from '@tanstack/react-query';
import { getCartOptions, getCartQueryKey } from '@/client/@tanstack/react-query.gen';

export default function TabsLayout() {
  const { colorScheme } = useColorScheme();
  const { data: cartItemsData } = useQuery({
    ...getCartOptions(),
    queryKey: getCartQueryKey(),
  });

  return (
    <Tabs
      screenOptions={{
        headerShown: false,
        tabBarActiveTintColor: colorScheme === 'dark' ? '#fff' : '#000',
        tabBarInactiveTintColor: colorScheme === 'dark' ? '#666' : '#999',
        tabBarStyle: {
          backgroundColor: colorScheme === 'dark' ? '#1a1a1a' : '#fff',
          borderTopColor: colorScheme === 'dark' ? '#333' : '#e5e5e5',
          borderTopWidth: 1,
          paddingBottom: 0,
          height: 60,
        },
      }}>
      <Tabs.Screen
        name="(home)"
        options={{
          title: 'Home',
          tabBarIcon: ({ color, size }) => <HomeIcon size={size} color={color} />,
        }}
      />
      <Tabs.Screen
        name="search"
        options={{
          title: 'Search',
          tabBarIcon: ({ color, size }) => <SearchIcon size={size} color={color} />,
        }}
      />
      <Tabs.Screen
        name="cart"
        options={{
          title: 'Cart',
          tabBarIcon: ({ color, size }) => <ShoppingCartIcon size={size} color={color} />,
          tabBarBadge: cartItemsData?.data?.items?.length ?? 0,

          tabBarBadgeStyle: {
            backgroundColor: 'red',
            color: 'white',
            fontSize: 12,
            fontWeight: 'bold',
            borderRadius: 10,
          },
        }}
      />
      <Tabs.Screen
        name="profile"
        options={{
          title: 'Profile',
          tabBarIcon: ({ color, size }) => <UserIcon size={size} color={color} />,
        }}
      />
    </Tabs>
  );
}
