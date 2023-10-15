<script>
	import { assets } from '$app/paths';
	import { page } from '$app/stores';
	import Viewer from '$lib/components/markdown/Viewer.svelte';

	let url = `${assets}/docs/${$page.params.section}/index.md`;
</script>

<Viewer {url} />