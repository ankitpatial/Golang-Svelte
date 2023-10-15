import {writable} from "svelte/store";

export interface VideoAccess {
  videoID: string;
  enabled: boolean;
}

export const videoAccess = writable<Array<VideoAccess>>([])
