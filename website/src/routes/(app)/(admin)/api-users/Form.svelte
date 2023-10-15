<script lang='ts'>
	import { AddApiUser } from '$lib/gql/apiUser';
	import { getContextClient } from '@urql/svelte';
	import Modal from '$lib/components/Modal.svelte';
	import Input from '$lib/components/form/Input.svelte';
	import alerts from '$lib/stores/alerts';
	import PwdField from './PwdField.svelte';

	export let reloadHandel: () => void;
	export let closeHandel: () => void;

	const client = getContextClient();
	let username = '', pwd = '', inp, busy = false;

	async function save() {
		if (busy || !username) return;

		const title = 'Api User';
		const res = await client.mutation(AddApiUser, { username }).toPromise();
		if (res.error) {
			alerts.push({ type: 'error', title, body: res.error.message });
			return;
		}

		pwd = res.data.addApiUser;
		alerts.push({ type: 'success', title, body: 'created successfully' });

		reloadHandel();
	}
</script>

<Modal open title='New Api User' cls='max-w-sm' okText='SAVE' on:cancel={closeHandel} {busy}>
	<form id='newApiUser' on:submit|preventDefault={save}>
		<Input
			name='username'
			label='API Username'
			placeholder='USERNAME'
			bind:value={username}
			required
			focus
		/>
		{#if (pwd)}
			<PwdField value={pwd} />
		{/if}
	</form>
	<svelte:fragment slot='actions'>
		<slot name='actions'>
			{#if (!pwd)}
				<button form='newApiUser' class='btn btn-primary' class:loading={busy} disabled={busy}>
					Save
				</button>
			{/if}
			<button class='btn btn-outline' on:click|preventDefault={closeHandel} disabled={busy}>
				{#if (pwd)}
					Close
				{:else}
					Cancel
				{/if}
			</button>
		</slot>
	</svelte:fragment>
</Modal>