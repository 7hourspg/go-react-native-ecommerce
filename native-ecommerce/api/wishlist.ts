import { useMutation, useQueryClient } from '@tanstack/react-query';
import {
  postWishlistMutation,
  getWishlistQueryKey,
  deleteWishlistByProductIdMutation,
  getWishlistByProductIdQueryKey,
} from '@/client/@tanstack/react-query.gen';
import { ModelsWishlistResponse } from '@/client/types.gen';

export const useAddToWishlist = () => {
  return useMutation({
    ...postWishlistMutation(),

    onMutate: async (newData, context) => {
      // Cancel outgoing refetches
      await context.client.cancelQueries({ queryKey: getWishlistQueryKey() });
      const checkQueryKey = getWishlistByProductIdQueryKey({
        path: { product_id: newData.body.product_id },
      });
      await context.client.cancelQueries({ queryKey: checkQueryKey });

      // Snapshot previous values
      const previousWishlist = context.client.getQueryData(getWishlistQueryKey());
      const previousCheck = context.client.getQueryData(checkQueryKey);

      // Optimistically update wishlist
      context.client.setQueryData(getWishlistQueryKey(), (old: ModelsWishlistResponse) => {
        if (!old?.data) return old;

        const tempItem = {
          ...newData.body,
          product: {
            id: newData.body.product_id,
            name: 'Loading...',
            description: 'Loading...',
            price: 0,
            image: '',
          },
        };

        return {
          ...old,
          data: {
            ...old.data,
            tempItem,
          },
        };
      });

      // Optimistically update check query
      context.client.setQueryData(checkQueryKey, {
        in_wishlist: true,
      });

      return { previousWishlist, previousCheck, checkQueryKey };
    },
    onError: (_, __, onMutateResult, context) => {
      context.client.setQueryData(getWishlistQueryKey(), onMutateResult?.previousWishlist);
    },
    onSettled: (_, __, ___, ____, context) => {
      context.client.invalidateQueries({ queryKey: getWishlistQueryKey() });
    },
  });
};

export const useRemoveFromWishlist = () => {
  return useMutation({
    ...deleteWishlistByProductIdMutation(),
    onMutate: async (newData, context) => {
      // Cancel outgoing refetches
      await context.client.cancelQueries({ queryKey: getWishlistQueryKey() });

      // Snapshot previous value
      const previousWishlist = context.client.getQueryData(getWishlistQueryKey());

      context.client.setQueryData(getWishlistQueryKey(), (old: ModelsWishlistResponse) => {
        return {
          ...old,
          data: old.data?.filter((item) => item.product_id !== newData.path.product_id),
        };
      });

      return { previousWishlist };
    },
    onError: (_, __, onMutateResult, context) => {
      context.client.setQueryData(getWishlistQueryKey(), onMutateResult?.previousWishlist);
    },
    onSettled: (_, __, ___, ____, context) => {
      context.client.invalidateQueries({ queryKey: getWishlistQueryKey() });
    },
  });
};
