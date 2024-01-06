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
			<a class="border-zinc-200 shadow-protruding rounded-md" href="/"
				><BxX /></a
			>
		</div>
		<div>
			{#await data}
				<p>Loading...</p>
			{:then data}
				<h3>{data.description}</h3>
				<p>Created {new Date(data.created_at).toLocaleString()}</p>
				<hr class="my-2 h-px border-0 bg-zinc-300" />
				<div class="flex justify-between mb-1">
					<h4 class="text-lg font-medium">Labels</h4>
					<a
						href="/project/{data.id}/label/create"
						class="block bg-teal-600 my-auto py-0.5 px-2 rounded-md border-blue-300 border text-white"
						>New Label</a
					>
				</div>
				{#if data.labels.length !== 0}
					<table
						class="w-full pt-2 border border-zinc-300 rounded-md border-spacing-1 text-center"
					>
						<thead>
							<tr class="border border-zinc-300">
								<th class="border border-zinc-300">ID</th>
								<th class="border border-zinc-300">Label</th>
								<th class="border border-zinc-300">Created At</th>
							</tr>
						</thead>
						<tbody>
							{#each data.labels as label}
								<tr class="border border-zinc-300">
									<td class="border border-zinc-300">
										<a
											class="text-teal-600"
											href="/project/{data.id}/label/{label.id}">{label.id}</a
										>
									</td>
									<td class="border border-zinc-300">{label.label}</td>
									<td class="border border-zinc-300"
										>{new Date(label.created_at).toLocaleString()}</td
									>
								</tr>
							{/each}
						</tbody>
					</table>
				{/if}
			{/await}
		</div>
	</div>
</div>
