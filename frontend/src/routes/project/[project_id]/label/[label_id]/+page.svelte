<script lang="ts">
	import type { ClassificationTaskLabel } from "$lib/types/label";
	import BxX from "svelte-boxicons/BxX.svelte";

	export let data: Promise<ClassificationTaskLabel>;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
		<div class="flex items-center justify-between">
			{#await data}
				<h2 class="page-title">Loading...</h2>
			{:then data}
				<h2 class="page-title">Label {data.label} (#{data.id})</h2>
			{/await}
			<a class="border-zinc-200 shadow-protruding rounded-md" href=".."
				><BxX /></a
			>
		</div>
		<div>
			{#await data}
				<p>Loading...</p>
			{:then data}
				<h3>Project {data.project_id}</h3>
				<p>Created {new Date(data.created_at).toLocaleString()}</p>
			{/await}
		</div>
	</div>
</div>
