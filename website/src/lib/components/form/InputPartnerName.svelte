<script lang='ts'>

	import Input from '$lib/components/form/Input.svelte';
	import { QryPartnerNameAvailable } from '$lib/gql/partner';
	import { getContextClient } from '@urql/svelte';

	export let label = 'Name';
	export let partnerType = '';
	export let partnerID = '';
	export let value = '';

	const client = getContextClient();

	async function coNameValidate(name) {
		if (!name || name.trim() == '') return;

		const res = await client.query(QryPartnerNameAvailable, { id: partnerID, type: partnerType, name }).toPromise();
		if (res.error) return res.error.message;

		const available = res.data?.partnerNameAvailable || false;
		return available ? '' : 'Name already in use';
	}
</script>

<Input {label} bind:value required maxlength='100' validate={coNameValidate}>
	<div slot='alt'><slot></slot></div>
</Input>