import { writable } from 'svelte/store';
import type { ProductInfo } from '$lib/graph/graphql';

export const editProduct = writable<ProductInfo | undefined>(undefined);
