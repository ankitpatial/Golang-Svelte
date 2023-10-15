<!--
 Copyright (c) 2022. Ankit Patial.
 Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 Author: Ankit Patial
  -->
<script lang='ts'>
	import { createEventDispatcher, onMount } from 'svelte';
	import Spin from '$lib/assets/svg/IconSpin.svelte';
	import imgNotFound from '$lib/assets/img/not-found.jpg';

	export let src;
	export let alt;
	export let cls;

	const dispatch = createEventDispatcher();
	let loaded = false;
	let failed = false;
	let loading = false;

	onMount(() => {
		const img = new Image();
		img.src = src;
		loading = true;

		img.onload = () => {
			loading = false;
			loaded = true;
			raiseDone();
		};

		img.onerror = () => {
			loading = false;
			failed = true;
			raiseDone();
		};
	});

	function raiseDone() {
		dispatch('done', { success: !failed });
	}
</script>

{#if loaded}
	<img {src} {alt} class='{cls}' />
{:else if failed}
	<img src='{imgNotFound}' class={cls} alt='Not Found' />
{:else if loading}
	<Spin {cls} />
{/if}
