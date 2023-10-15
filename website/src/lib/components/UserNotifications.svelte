<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
	import { getContextClient } from '@urql/svelte';
	import alerts from '$lib/stores/alerts';
	import { humanize } from '$lib/utils/string.js';
	import InputToggle from '$lib/components/form/InputToggle.svelte';
	import { QryUserNotifySettings, SaveNotifySettings } from '$lib/gql/user';

	export let userID;

	const client = getContextClient();

	let busy, data;


	// get user's possible notifications
	$: if (userID) {
		busy = true;
		client.query(QryUserNotifySettings, { id: userID }).toPromise().then((res) => {
			if (res.error) {
				alerts.error('Notification Settings', res.error.message);
				return;
			}
			data = res.data.userNotifySettings;
		}).finally(() => {
			busy = false;
		});
	}

	// save a user notification flag value
	async function save(topicID: string, val: boolean) {
		const res = await client.mutation(SaveNotifySettings, { id: userID, topicID, email: val }).toPromise();
		if (res.error) {
			alerts.error('Notification Setting Saved Successfully', res.error.message);
			return;
		}
	}
</script>

<div class:blur-sm={busy}>
	<table class='table table-compact'>
		<thead>
		<tr>
			<td>Topic</td>
			<td>Receive Email</td>
		</tr>
		</thead>
		<tbody>
		{#each data || [] as item}
			<tr>
				<td>{humanize(item.topic)}</td>
				<td>
					<InputToggle value={item.receiveEmail} on:change={({detail}) => save(item.id, detail)} />
				</td>
			</tr>
		{/each}
		</tbody>
	</table>
</div>
