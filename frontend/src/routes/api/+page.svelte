<script lang="ts">
	import { onMount } from 'svelte';
	import SwaggerUI from 'swagger-ui';
	import 'swagger-ui/dist/swagger-ui.css';

	onMount(async () => {
	 SwaggerUI({
		url: '/api/openapi.json',
		dom_id: '#swagger-ui-container'
	 });
	});
 </script>

 <svelte:head>
	<title>Semantic Sensei API</title>
 </svelte:head>

 <div id="swagger-ui-container" />
