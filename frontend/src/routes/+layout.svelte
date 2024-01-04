<script>
	import "./+layout.scss";
</script>

<div class="w-screen max-w-full min-h-screen flex flex-col">
	<div class="sticky top-0 shadow-lg border-md flex px-3 py-2 bg-zinc-50">
		<div class="flex-initial flex flex-col items-center">
			<h1 class="text-2xl">
				<a href="/">Semantic Sensei</a>
			</h1>
		</div>
		<div class="flex-initial ml-auto flex flex-col items-center">
			<a class="block bg-blue-500 my-auto py-1 px-2 rounded-md border-blue-300 border text-white" href="/project/create">New Project</a>
		</div>
	</div>
	<div class="flex-1 flex flex-col items-center">
		<slot />
	</div>
</div>
