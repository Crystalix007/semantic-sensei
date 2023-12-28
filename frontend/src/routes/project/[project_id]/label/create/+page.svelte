<script lang="ts">
	import type { PageProps } from "./+page";

	export let data: PageProps;
</script>

<div class="centred-page">
	<div class="h-fit w-fit border-2 rounded-lg p-6 flex flex-col gap-4">
		<h1>Create a new label</h1>
		<form class="flex flex-col gap-2" action={data.form_url} method="post">
			<div class="flex flex-col gap-1">
				<label for="name">Label</label>
				<input type="text" id="label" name="label" required />
			</div>
			<button class="create-button" type="submit">Create</button>
		</form>
	</div>
</div>
