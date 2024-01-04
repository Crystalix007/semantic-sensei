<script lang="ts">
	import type { Project } from "$lib/types/project";
	import BxX from "svelte-boxicons/BxX.svelte";

	export let data: Promise<Project>;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
		<div class="flex items-center justify-between">
			{#await data}
				<h2 class="page-title">Loading...</h2>
			{:then data}
				<h2 class="page-title">Project {data.name} (#{data.id})</h2>
			{/await}
			<a class="border-zinc-200 shadow-protruding rounded-md" href="/"><BxX /></a>
		</div>
		<div>
			{#await data}
				<p>Loading...</p>
			{:then data}
				<h3>{data.description}</h3>
				<p>Created {new Date(data.created_at).toLocaleString()}</p>
			{/await}
		</div>
	</div>
</div>
