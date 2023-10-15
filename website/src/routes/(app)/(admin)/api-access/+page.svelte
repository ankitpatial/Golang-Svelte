<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import { QryApiAccess } from '$lib/gql/apiAccess';
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import TableCard from '$lib/components/TableCard.svelte';
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { page } from '$app/stores';
  import Page from '$lib/components/Page.svelte';
  import UpdateApiAccessKey from './UpdateApiAccessKey.svelte';
  import UpdateApiAccessSecret from './UpdateApiAccessSecret.svelte';

  $: variables = {
    search: $page.url.searchParams.get('q')
  };

  function add() {
    goto(PATH.API_ACCESS_CREATE);
  }
</script>

<Page title='API Access'>
  <svelte:fragment slot='buttons'>
    <BtnPlusSmall on:click={add} />
  </svelte:fragment>

  <TableCard>
    <DataTable
      headers={['NAME', 'URL', 'USERNAME', 'ACTIONS' ]}
      query={QryApiAccess}
      dataProp='apiAccess'
      pageSize='10'
      searchable
      {variables}
    >
      <svelte:fragment slot='row' let:headers let:item>
        <DataColumn cls='md:w-96' header={headers[0]}>{item.id}</DataColumn>
        <DataColumn cls='md:w-96' header={headers[1]}>{item.url}</DataColumn>
        <DataColumn cls='md:w-96' header={headers[2]}>{item.username}</DataColumn>
        <DataColumn cls='md:w-96' header={headers[2]}>
          <UpdateApiAccessKey id={item.id} />
        </DataColumn>
      </svelte:fragment>
    </DataTable>
  </TableCard>
</Page>
