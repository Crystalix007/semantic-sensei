<script lang="ts">
	import BxX from "svelte-boxicons/BxX.svelte";
	import type { PageProps } from "./+page";

	export let data: PageProps;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
		<div class="flex items-center justify-between">
			<h2 class="page-title">Create a new task</h2>
			<a class="border-zinc-200 shadow-protruding rounded-md" href=".."
				><BxX /></a
			>
		</div>
		<form class="flex flex-col gap-2" action={data.form_url} method="post">
			<div class="flex flex-col gap-1">
				<label for="llm_input">Input</label>
				<textarea id="llm_input" name="llm_input" required />
				<label for="llm_output">Output</label>
				<textarea id="llm_output" name="llm_output" required />
				<label for="embedding">Embedding</label>
				<textarea id="embedding" name="embedding" required />
			</div>
			<button class="create-button" type="submit">Create</button>
		</form>
	</div>
</div>
