<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { PATH } from '$lib/enum.js';
  import { QryApiUserAuditLogs } from '$lib/gql/audit';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import TableCard from '$lib/components/TableCard.svelte';
  import { DataColumn, DataTable } from '$lib/components/DataTable/index.js';
  import Page from '$lib/components/Page.svelte';

  $: variables = {
    id: $page.params['id'],
    search: $page.url.searchParams.get('q'),
    orderBy: {
      field: 'CREATED',
      direction: 'DESC'
    }
  };

  function back() {
    goto(PATH.API_USERS);
  }
</script>

<Page title='Audit Logs'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>

  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <TableCard>
        <DataTable
          headers={['Date', 'Action', 'Description', 'IP']}
          query={QryApiUserAuditLogs}
          dataProp='apiUserAuditLogs'
          pageSize='10'
          searchable
          {variables}
        >
          <svelte:fragment slot='row' let:headers let:item>
            <DataColumn header={headers[0]} cls='w-52'>
              {item.createdAt}
            </DataColumn>
            <DataColumn header={headers[1]} cls='w-52'>
              {item.action}
            </DataColumn>
            <DataColumn header={headers[2]}>
              {item.description}
            </DataColumn>
            <DataColumn header={headers[3]} cls='w-52'>
              {item.ip}
            </DataColumn>
          </svelte:fragment>
        </DataTable>
      </TableCard>
    </div>
  </div>
</Page>
