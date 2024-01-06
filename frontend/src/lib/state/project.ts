import { writable } from "svelte/store";

export const project = writable<string | undefined>(undefined);
