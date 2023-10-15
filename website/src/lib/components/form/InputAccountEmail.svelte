<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  -->

<script lang='ts'>
	import { getContextClient } from '@urql/svelte';
	import { QryUserEmailAvailable } from '$lib/gql/account';
	import Input from '$lib/components/form/Input.svelte';

	export let userID = '';
	export let value = '';
	export let label = 'Email';
	export let err = '';
  export let handleChange: (event: Event) => any = null;

  const client = getContextClient();

	// validate user email
	async function userEmailValidate(userID, email) {
		if (!email || email.trim() == '') return '';

		const res = await client.query(QryUserEmailAvailable, { id: userID || '', email }).toPromise();
		if (res.error) return res.error.message;

		return res.data.userEmailAvailable ? '' : 'Already in use';
	}
</script>

<Input
	{label}
	maxlength='100'
	bind:value={value}
	validate={(v) => userEmailValidate(userID, v)}
	required
  {err}
  {handleChange}
/>
