<script lang="ts">
	import type { PageProps } from "./+page";
	import BxX from "svelte-boxicons/BxX.svelte";

	export let data: Promise<PageProps>;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4 m-8">
		<div class="flex items-center justify-between">
			{#await data}
				<h2 class="page-title">Loading...</h2>
			{:then data}
				<h2 class="page-title">Task {data.task.id}</h2>
			{/await}
			<a class="border-zinc-200 shadow-protruding rounded-md" href=".."
				><BxX /></a
			>
		</div>
		<div>
			{#await data}
				<p>Loading...</p>
			{:then data}
				<p>Created {new Date(data.task.created_at).toLocaleString()}</p>
				<hr class="h-px border-0 bg-slate-300 my-2" />
				<div class="flex justify-between gap-4">
					<div>
						<div class="text-lg font-semibold pb-1">Input</div>
						<pre class="whitespace-pre-wrap">{data.task.llm_input}</pre>
					</div>
					<div>
						<div class="text-lg font-semibold pb-1">Output</div>
						<pre class="whitespace-pre-wrap">{data.task.llm_output}</pre>
					</div>
				</div>
				<hr class="h-px border-0 bg-slate-300 mt-3 mb-2" />
				<form
					class="w-full flex flex-wrap justify-end gap-1"
					action={data.label_endpoint}
					method="post"
				>
					{#each data.labels as label}
						<button class="create-button" name="label" value={label.id}
							>{label.label}</button
						>
					{/each}
				</form>
			{/await}
		</div>
	</div>
</div>
