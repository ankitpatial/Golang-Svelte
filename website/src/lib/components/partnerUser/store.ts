import {writable} from "svelte/store";
import type {PartnerContact} from "$lib/graph/graphql";

export const editUser = writable<PartnerContact>()
