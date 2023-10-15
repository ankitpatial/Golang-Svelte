<script lang='ts'>
	import Input from '$lib/components/form/Input.svelte';
	import { copyToClipboard } from '$lib/utils/clipboard';
	import IconDuplicate from '$lib/assets/svg/IconDuplicate.svelte';


	export let value;
	let inp;
</script>

<Input bind:ref={inp} label='Password' {value}>
	<svelte:fragment slot='group'>
		<button class='btn btn-primary' on:click|preventDefault={copyToClipboard(inp)}>
			<IconDuplicate />
		</button>
	</svelte:fragment>
</Input>
<span><i class='text-gray-400'>please copy this password before you close this modal</i></span>