import { derived, writable } from 'svelte/store';

export const resolution = writable<number>(600);

export const smallScreen = derived(resolution, ($res) => {
  return $res < 640;
});
