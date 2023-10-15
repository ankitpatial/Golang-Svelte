<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
	import { goto } from '$app/navigation';
	import { API_ACCESS_SOURCE, PATH } from '$lib/enum';
	import type ApiAccess from '$lib/models/ApiAccess';
	import Input from '$lib/components/form/Input.svelte';
	import DropDown from '$lib/components/form/DropDown.svelte';
	import { getContextClient } from '@urql/svelte';
	import { SaveApiAccess } from '$lib/gql/apiAccess';
	import alerts from '$lib/stores/alerts';

	export let data: ApiAccess = {
		id: '',
		name: API_ACCESS_SOURCE.EAGLE_VIEW,
		username: '',
		password: '',
		url: '',
		key: '',
		secret: ''
	};


	const client = getContextClient();
	let busy = false;

	async function save() {
		if (busy) return;

		busy = true;
		const res = await client.mutation(SaveApiAccess, { input: data }).toPromise();
		busy = false;

		const title = 'Api Access Info';
		if (res.error) {
			alerts.push({ type: 'error', title, body: res.error.message });
			return;
		}

		alerts.push({ type: 'success', title, body: 'successfully saved' });
		await goto(PATH.API_ACCESS);
	}

</script>

<div class='w-full md:w-96'>
	<form id='newApiUser' on:submit|preventDefault={save}>
		<DropDown
			label='Name'
			name='name'
			bind:value={data.name}
			options={[
				{id: '',  name: 'Other'},
				{id: API_ACCESS_SOURCE.EAGLE_VIEW,  name: 'Eagle View'},
				{id: API_ACCESS_SOURCE.NEAR_MAP,  name: 'Nearmap'},
				{id: API_ACCESS_SOURCE.REGRID,  name: 'ReGrid'}
			]}
			required
		/>

		<Input
			label='Url'
			name='url'
			type='url'
			placeholder='name'
			bind:value={data.url}
			required
		/>

		<Input
			label='API Username'
			name='username'
			placeholder='username'
			bind:value={data.username}
			required
		/>


		<Input
			label='API Password'
			name='password'
			placeholder='password'
			bind:value={data.password}
			required
		/>

		<Input
			label='Key'
			name='key'
			placeholder='key, source id etc.'
			bind:value={data.key}
			required
		/>

		<Input
			label='Secret'
			name='secret'
			placeholder='secret, client secret etc.'
			bind:value={data.secret}
			required
		/>

		<button type='submit' class='btn btn-primary' class:loading={busy}>
			Save
		</button>
	</form>
</div>
