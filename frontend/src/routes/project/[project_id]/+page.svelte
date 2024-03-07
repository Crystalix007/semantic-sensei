<script lang="ts">
	import BxX from "svelte-boxicons/BxX.svelte";
	import type { PageData } from "./+page";

	export let data: Promise<PageData>;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
		<div class="flex items-center justify-between">
			{#await data}
				<h2 class="page-title">Loading...</h2>
			{:then data}
				<h2 class="page-title">
					Project {data.project.name} (#{data.project.id})
				</h2>
			{/await}
			<a class="ml-1 border-zinc-200 shadow-protruding rounded-md" href="/"
				><BxX /></a
			>
		</div>
		<div>
			{#await data}
				<p>Loading...</p>
			{:then data}
				<h3>{data.project.description}</h3>
				<p>Created {new Date(data.project.created_at).toLocaleString()}</p>
				<hr class="my-2 h-px border-0 bg-zinc-300" />
				<div class="flex justify-between mb-1">
					<h4 class="text-lg font-medium">Labels</h4>
					<a
						href="/project/{data.project.id}/label/create"
						class="block bg-teal-600 my-auto py-0.5 px-2 rounded-md border-blue-300 border text-white"
						>New Label</a
					>
				</div>
				{#if data.project.labels.length !== 0}
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
							{#each data.project.labels as label}
								<tr class="border border-zinc-300">
									<td class="border border-zinc-300">
										<a
											class="text-teal-600"
											href="/project/{data.project.id}/label/{label.id}"
											>{label.id}</a
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
				<hr class="my-2 h-px border-0 bg-zinc-300" />
				<div class="flex justify-between mb-1">
					<h4 class="text-lg font-medium">Pending Tasks</h4>
					<a
						href="/project/{data.project.id}/task/create"
						class="block bg-teal-600 my-auto py-0.5 px-2 rounded-md border-blue-300 border text-white"
						>New Task</a
					>
				</div>
				{#if data.pending_tasks.data.length !== 0}
					<table
						class="w-full pt-2 border border-zinc-300 rounded-md border-spacing-1 text-center"
					>
						<thead>
							<tr class="border border-zinc-300">
								<th class="border border-zinc-300">ID</th>
								<th class="border border-zinc-300">Created At</th>
							</tr>
						</thead>
						<tbody>
							{#each data.pending_tasks.data as task}
								<tr class="border border-zinc-300">
									<td class="border border-zinc-300">
										<a
											class="text-teal-600"
											href="/project/{data.project.id}/task/{task.id}"
											>{task.id}</a
										>
									</td>
									<td class="border border-zinc-300"
										>{new Date(task.created_at).toLocaleString()}</td
									>
								</tr>
							{/each}
						</tbody>
					</table>
				{/if}
				{#if data.completed_tasks.data.length !== 0}
					<hr class="my-2 h-px border-0 bg-zinc-300" />
					<h4 class="text-lg font-medium mb-1">Completed Tasks</h4>
					<table
						class="w-full pt-2 border border-zinc-300 rounded-md border-spacing-1 text-center"
					>
						<thead>
							<tr class="border border-zinc-300">
								<th class="border border-zinc-300">ID</th>
								<th class="border border-zinc-300">Created At</th>
								<th class="border border-zinc-300">Label</th>
							</tr>
						</thead>
						<tbody>
							{#each data.completed_tasks.data as task}
								<tr class="border border-zinc-300">
									<td class="border border-zinc-300">
										<a
											class="text-teal-600"
											href="/project/{data.project.id}/task/{task.id}"
											>{task.id}</a
										>
									</td>
									<td class="border border-zinc-300"
										>{new Date(task.created_at).toLocaleString()}</td
									>
									<td class="border border-zinc-300">
										{#if task.label_id}
											<a
												class="text-teal-700"
												href="/project/{data.project.id}/label/{task.label_id}"
											>
												{task.label_id}
											</a>
										{/if}
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				{/if}
			{/await}
		</div>
	</div>
</div>
