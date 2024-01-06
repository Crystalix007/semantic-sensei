<script lang="ts">
	import BxX from "svelte-boxicons/BxX.svelte";
	import type { PageProps } from "./+page";

	export let data: PageProps;
</script>

<div class="top-page">
	<div class="border-2 rounded-lg p-6 flex flex-col gap-4">
		<div class="flex items-center justify-between">
			<h2 class="page-title">Create a new label</h2>
			<a class="border-zinc-200 shadow-protruding rounded-md" href=".."
				><BxX /></a
			>
		</div>
		<form class="flex flex-col gap-2" action={data.form_url} method="post">
			<div class="flex flex-col gap-1">
				<label for="name">Label</label>
				<input type="text" id="label" name="label" required />
			</div>
			<button class="create-button" type="submit">Create</button>
		</form>
	</div>
</div>
