<script>
	import { assets } from '$app/paths';
	import Viewer from '$lib/components/markdown/Viewer.svelte';

	let url = `${assets}/docs/index.md`;
</script>

<svelte:head>
	<title>API Documentation</title>
</svelte:head>

<Viewer {url} />