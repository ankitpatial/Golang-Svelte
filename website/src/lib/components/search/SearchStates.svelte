<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
	import { getContextClient } from '@urql/svelte';
	import { QryStates } from '$lib/gql/postalCode';
	import SearchBox from '$lib/components/form/SearchBox.svelte';
	import alerts from '$lib/stores/alerts';

	export let label = '';
	export let placeholder;
	export let cls = '';

	const client = getContextClient();
	const title = 'Search Cities';

	let open = false;

	async function search(q: string) {
		const res = await client.query(QryStates, { q: q.trim(), skip: 0, take: 5 }).toPromise();
		if (res.error) {
			alerts.push({ type: 'error', title: `Error: Search State`, body: res.error.message });
			return [];
		}

		return res.data.states;
	}

</script>

<SearchBox
	{label}
	{placeholder}
	{cls}
	searchHandler={search}
	noSelection
	on:change
/>