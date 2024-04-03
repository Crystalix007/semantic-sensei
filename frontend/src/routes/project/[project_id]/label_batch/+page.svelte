<script lang="ts">
	import type { PageProps } from "./+page";
	import BxX from "svelte-boxicons/BxX.svelte";

	export let data: Promise<PageProps>;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
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
				<div class="flex flex-col justify-between gap-4">
					<details open class="hide-marker">
						<summary class="text-lg font-semibold pb-1">Input</summary>
						<pre class="whitespace-pre-wrap">{data.task.llm_input}</pre>
					</details>
					<div>
						<div class="text-lg font-semibold pb-1">Output</div>
						<pre class="whitespace-pre-wrap">{data.task.llm_output}</pre>
					</div>
				</div>
				<hr class="h-px border-0 bg-slate-300 mt-3 mb-2" />
				<div class="w-full sticky bottom-4">
					<form
						class="ml-auto w-fit flex flex-wrap justify-between md:justify-end gap-2 md:gap-3 bg-white border-slate-300 border-2 b-1 p-2 md:p-3 rounded-md"
						action={data.label_endpoint}
						method="post"
					>
						{#each data.labels as label}
							<button class="create-button" name="label" value={label.id}
								>{label.label}</button
							>
						{/each}
					</form>
				</div>
			{/await}
		</div>
	</div>
</div>

<style lang="postcss">
	@media (max-width: theme("screens.md")) {
		details.hide-marker > summary:first-of-type::marker {
			@apply transition-all;
		}

		details.hide-marker > summary:first-of-type::-webkit-details-marker {
			@apply transition-all;
		}

		details.hide-marker > summary:first-of-type {
			@apply cursor-pointer;
		}

		details[open].hide-marker > summary:first-of-type {
			@apply list-none;
		}

		details[open].hide-marker > summary:first-of-type::-moz-list-bullet {
			@apply list-none block text-transparent w-0 m-0;
		}

		details[open].hide-marker > summary:first-of-type::-webkit-details-marker {
			@apply invisible w-0 m-0;
		}
	}

	@screen md {
		details.hide-marker > summary:first-of-type {
			@apply list-none pointer-events-none;
		}
	}
</style>
