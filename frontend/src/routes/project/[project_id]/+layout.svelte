<script lang="ts">
	import { project } from "$lib/state/project.js";
	import type { Project } from "$lib/types/project";
	import { onDestroy, onMount } from "svelte";

	export let data: Promise<Project>;

	onDestroy(() => {
		project.set(undefined);
	});

	onMount(async () => {
		const projectData = await data;

		project.set(projectData.name);
	});
</script>

<slot />
