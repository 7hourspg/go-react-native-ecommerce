import { useMutation, useQueryClient } from '@tanstack/react-query';
import {
  postCartItemsMutation,
  putCartItemsByIdMutation,
  getCartQueryKey,
  deleteCartItemsByIdMutation,
} from '@/client/@tanstack/react-query.gen';
import { ModelsCartSummaryResponse } from '@/client/types.gen';

export const useAddToCart = () => {
  const queryClient = useQueryClient();
  return useMutation({
    ...postCartItemsMutation(),

    onMutate: async (newData) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: getCartQueryKey() });

      // Snapshot previous value
      const previousCart = queryClient.getQueryData(getCartQueryKey());

      queryClient.setQueryData(getCartQueryKey(), (old: ModelsCartSummaryResponse) => {
        if (!old?.data) return old;

        const existingItems = old.data.items ?? [];

        // Check if item already exists in cart
        const existingItemIndex = existingItems.findIndex(
          (item) => item.product_id === newData.body.product_id
        );

        let updatedItems: typeof existingItems;

        if (existingItemIndex >= 0) {
          // Item exists - increase quantity
          updatedItems = existingItems.map((item, index) =>
            index === existingItemIndex
              ? { ...item, quantity: (item.quantity ?? 0) + newData.body.quantity }
              : item
          );
        } else {
          // Item doesn't exist - add new item
          const tempItem = {
            ...newData.body,
            quantity: newData.body.quantity,
            product: {
              id: newData.body.product_id,
              name: 'Loading...',
              description: 'Loading...',
              price: 0,
              image: '',
            },
          };
          updatedItems = [...existingItems, tempItem];
        }

        return {
          ...old,
          data: {
            ...old.data,
            items: updatedItems,
          },
        };
      });

      return { previousCart };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(getCartQueryKey(), context?.previousCart);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: getCartQueryKey() });
    },
  });
};

export const useUpdateCartItemQuantity = () => {
  const queryClient = useQueryClient();

  return useMutation({
    ...putCartItemsByIdMutation(),
    onMutate: async (newData) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: getCartQueryKey() });

      // Snapshot previous value
      const previousCart = queryClient.getQueryData(getCartQueryKey());

      queryClient.setQueryData(getCartQueryKey(), (old: ModelsCartSummaryResponse) => {
        return {
          ...old,
          data: {
            ...old.data,
            items: old.data?.items?.map((item) =>
              item.id === newData.path.id ? { ...item, quantity: newData.body.quantity } : item
            ),
            isLoading: true,
          },
        };
      });

      return { previousCart };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(getCartQueryKey(), context?.previousCart);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: getCartQueryKey() });
      queryClient.setQueryData(getCartQueryKey(), (old: ModelsCartSummaryResponse) => {
        return {
          ...old,
          data: {
            ...old.data,
            isLoading: false,
          },
        };
      });
    },
  });
};

export const useRemoveFromCart = () => {
  const queryClient = useQueryClient();
  return useMutation({
    ...deleteCartItemsByIdMutation(),
    onMutate: async (newData) => {
      // Cancel outgoing refetches
      await queryClient.cancelQueries({ queryKey: getCartQueryKey() });

      // Snapshot previous value
      const previousCart = queryClient.getQueryData(getCartQueryKey());

      queryClient.setQueryData(getCartQueryKey(), (old: ModelsCartSummaryResponse) => {
        return {
          ...old,
          data: {
            ...old.data,
            items: old.data?.items?.filter((item) => item.id !== newData.path.id),
            isLoading: true,
          },
        };
      });

      return { previousCart };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData(getCartQueryKey(), context?.previousCart);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: getCartQueryKey() });
      queryClient.setQueryData(getCartQueryKey(), (old: ModelsCartSummaryResponse) => {
        return {
          ...old,
          data: {
            ...old.data,
            isLoading: false,
          },
        };
      });
    },
  });
};
