<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {QryPartners} from '$lib/gql/partner';
  import SearchBox from '$lib/components/form/SearchBox.svelte';
  import alerts from '$lib/stores/alerts';
  import type Entity from '$lib/models/Entity';

  export let partnerType;
  export let value: Entity;

  const client = getContextClient();

  async function search(q: string) {
    const res = await client.query(QryPartners, {search: q, partnerType, page: {first: 10}}).toPromise();
    if (res.error) {
      alerts.push({type: 'error', title: 'Get Users', body: res.error.message});
      return [];
    }

    return res.data.partners.edges.map(({node}) => ({id: node.id, name: node.name}));
  }
</script>

<SearchBox label='Partner' searchHandler={search} bind:value/>
