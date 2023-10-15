<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script>
	import { authUser } from '$lib/stores/auth';
	import { PATH } from '$lib/enum';
	import { goto } from '$app/navigation';
	import IconSignOut from '$lib/assets/svg/IconSignOut.svelte';

	export let cls = '';
	export let label = 'Sign Out';

	async function signOut() {
		await fetch(PATH.LOGOUT, { method: 'DELETE' });
		authUser.set(null);
		await goto(PATH.LOGIN);
	}

	$: user = $authUser;
</script>

<a on:click={signOut} class={cls}>
	<IconSignOut /> {label}
</a>