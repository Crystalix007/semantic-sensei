<script lang="ts">
	import type { Projects } from "$lib/types/projects";

	export let data: Promise<Projects>;
</script>

<div class="top-page m-8">
	<h2 class="text-4xl">Semantic Sensei - Build Your Own Semantic Classifier</h2>
	<p class="pt-2">
		Manually label samples of LLM text output, and train a classifier to label
		it automatically.
	</p>
	<hr class="my-2 h-px border-0 bg-zinc-300" />
	{#await data}
		<p class="pt-2">Loading...</p>
	{:then data}
		<h4 class="text-lg font-medium mb-1">Projects</h4>
		<table
			class="w-full border border-zinc-300 rounded-md border-spacing-1 text-center"
		>
			<thead>
				<tr class="border border-zinc-300">
					<th class="border border-zinc-300">Project</th>
					<th class="border border-zinc-300">Description</th>
				</tr>
			</thead>
			<tbody>
				{#each data.projects as project}
					<tr class="border border-zinc-300">
						<td class="border border-zinc-300">
							<a class="text-teal-600" href="project/{project.id}"
								>{project.name}</a
							>
						</td>
						<td class="border border-zinc-300">{project.description}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/await}
</div>
