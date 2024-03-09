<script lang="ts">
	import "./+layout.scss";
	import { project } from "$lib/state/project";
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
	<div class="sticky top-0 shadow-lg border-md flex px-3 md:px-8 py-2 bg-zinc-50">
		<div class="flex-initial flex flex-col items-center">
			<h1 class="text-2xl">
				<a href="/">Semantic Sensei</a>
			</h1>
		</div>

		{#if $project}
			<div class="flex-1 flex flex-col items-center">
				<h2 class="text-xl my-auto font-semibold">{$project}</h2>
			</div>
		{/if}

		<div class="flex-initial ml-auto flex gap-2 items-center">
			<a
				class="h-full block bg-blue-500 my-auto py-1 px-2 rounded-md border-blue-300 border text-white"
				href="/project/create">New Project</a
			>
		</div>
	</div>
	<div class="flex-1 flex flex-col items-center">
		<slot />
	</div>
</div>
