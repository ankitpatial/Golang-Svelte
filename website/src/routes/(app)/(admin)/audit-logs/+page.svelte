<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { QryAuditLogs } from '$lib/gql/audit';
  import TableCard from '$lib/components/TableCard.svelte';
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { page } from '$app/stores';
  import Page from '$lib/components/Page.svelte';

  $: variables = {
    search: $page.url.searchParams.get('q'),
    orderBy: {
      field: 'CREATED',
      direction: 'DESC'
    }
  };
</script>

<Page title='Audit Logs'>
  <TableCard>
    <DataTable
      headers={['Date', 'Action', 'Description', 'USER', 'IP']}
      query={QryAuditLogs}
      dataProp='auditLogs'
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

        <DataColumn header={headers[4]} cls='w-32'>
          {item.user ? `${item.user.firstName} ${item.user.lastName}` : ''}
          {item.apiUser ? `API: ${item.apiUser.username}` : ''}
        </DataColumn>
        <DataColumn header={headers[3]} cls='w-32'>
          {item.ip}
        </DataColumn>
      </svelte:fragment>
    </DataTable>
  </TableCard>
</Page>
