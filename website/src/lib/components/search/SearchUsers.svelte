<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { QryUsersSearch } from '$lib/gql/user';
  import SearchBox from '$lib/components/form/SearchBox.svelte';
  import alerts from '$lib/stores/alerts';
  import type Entity from '$lib/models/Entity';
  import { createEventDispatcher } from 'svelte';

  export let value: Entity;

  const client = getContextClient();
  const dispatch = createEventDispatcher();

  async function search(q: string) {
    const res = await client.query(QryUsersSearch, { search: q }).toPromise();
    if (res.error) {
      alerts.push({ type: 'error', title: 'Get Users', body: res.error.message });
      return [];
    }

    return res.data.usersSearch;
  }

  function change({detail}) {
    dispatch('change', detail)
  }
</script>

<SearchBox
  label='Search User'
  placeholder='type name or email'
  searchHandler={search}
  bind:value
  render={(i) => `${i.firstName} ${i.lastName} (${i.partnerName || 'RFX' })`}
  on:change={change}
/>
