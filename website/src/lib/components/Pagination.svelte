<script lang='ts'>
	import { createEventDispatcher } from 'svelte';

	export let hasPrevious = false;
	export let hasNext = false;

	const dispatch = createEventDispatcher();

	function previousPage() {
		dispatch('previous');
	}

	function nextPage() {
		dispatch('next');
	}
</script>
<div class='btn-group justify-end'>
	<button class='btn btn-outline btn-sm' on:click={previousPage} disabled={!hasPrevious}>
		Previous
	</button>
	<button class='btn btn-outline btn-sm' on:click={nextPage} disabled={!hasNext}>
		Next
	</button>
</div>