<script>
	import { assets } from '$app/paths';
	import { page } from '$app/stores';
	import Viewer from '$lib/components/markdown/Viewer.svelte';

	let url = '';
	$: url = `${assets}/docs/${$page.params.section}/${$page.params.file}.md`;
</script>

<Viewer {url} />





