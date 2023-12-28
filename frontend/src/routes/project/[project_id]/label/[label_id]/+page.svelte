<script lang="ts">
	import type { ClassificationTaskLabel } from "$lib/types/label";

	export let data: Promise<ClassificationTaskLabel>;
</script>

<div class="centred-page">
	<div>
		{#await data}
			<p>Loading...</p>
		{:then data}
			<h1>Label {data.label} (#{data.id})</h1>
			<h2>Project {data.project_id}</h2>
			<p>Created {new Date(data.created_at).toLocaleString()}</p>
		{/await}
	</div>
</div>
