import { writable } from 'svelte/store';

export const online =writable<boolean>(false);
export const unreadNotificationsCount = writable<number>(0);
