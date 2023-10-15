<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
	import { getContextClient } from '@urql/svelte';
	import { QryCities } from '$lib/gql/postalCode';
	import SearchBox from '$lib/components/form/SearchBox.svelte';
	import alerts from '$lib/stores/alerts';
	import type Entity from '$lib/models/Entity';

	export let state: string;
	export let label = '';
	export let placeholder;
	export let cls = '';
	export let ignore: Array<Entity> = [];

	const client = getContextClient();
	const title = 'Search Cities';

	let open = false;

	async function search(q: string) {
		const res = await client.query(QryCities, { state, q: q.trim(), skip: 0, take: 10 }).toPromise();
		if (res.error) {
			alerts.push({ type: 'error', title, body: res.error.message });
			return [];
		}

		return res.data.cities;
	}

</script>

<SearchBox
	{label}
	{placeholder}
	{cls}
	value={ignore}
	searchHandler={search}
	noSelection
	on:change
	render={(i) => `${i.name} - ${i.zip}`}
/>