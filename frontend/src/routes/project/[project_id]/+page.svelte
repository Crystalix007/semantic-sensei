<script lang="ts">
	import type { Project } from "$lib/types/project";

	export let data: Promise<Project>;
</script>

<div class="centred-page">
	<div>
		{#await data}
			<p>Loading...</p>
		{:then data}
			<h1>Project {data.name} (#{data.id})</h1>
			<h2>{data.description}</h2>
			<p>Created {new Date(data.created_at).toLocaleString()}</p>
		{/await}
	</div>
</div>
