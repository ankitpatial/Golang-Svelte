import { writable } from 'svelte/store';
import type { ProductPackage } from '$lib/graph/graphql';


export const editPackage = writable<ProductPackage | undefined>(undefined);
