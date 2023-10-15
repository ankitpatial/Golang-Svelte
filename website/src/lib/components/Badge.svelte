<script lang='ts'>
	export let type: 'success' | 'warning' | 'error' | 'outline' = 'outline';
</script>

<span class='badge badge-{type} border-0 py-0.5 px-2 text-xs text-center'>
	<slot />
</span>