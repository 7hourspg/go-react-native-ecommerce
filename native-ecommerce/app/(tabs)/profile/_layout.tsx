// app/(tabs)/profile/_layout.tsx
// This layout is used to display the profile screen, wishlist screen, and orders screen so the route will be like '(tabs)/profile/wishlist' or '(tabs)/profile/orders'
import { Stack } from 'expo-router';

export default function ProfileLayout() {
  return (
    <Stack screenOptions={{ headerShown: false }}>
      <Stack.Screen name="index" />
      <Stack.Screen name="wishlist" />
      <Stack.Screen name="orders" />
    </Stack>
  );
}
