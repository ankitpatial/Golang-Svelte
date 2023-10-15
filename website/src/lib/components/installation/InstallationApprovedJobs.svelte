<script lang='ts'>
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import { InstallationType } from '$lib/graph/graphql';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { usdFormat } from '$lib/utils/currency';
  import { localDate, thisMonth } from '$lib/utils/time';
  import { authUser } from '$lib/stores/auth';
  import PendingJobActions from './PendingJobActions.svelte';
  import Page from '$lib/components/Page.svelte';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import { page } from '$app/stores';
  import { QryApprovedInstallations } from '$lib/gql/installation.js';
  import DDLStatus from '$lib/components/installation/DDLStatus.svelte';

  export let title: string;
  export let installationType: InstallationType;
  export let basePath;

  const isAdmin = $authUser?.isAdmin || $authUser?.isCompanyAdmin || false;
  const headers = [
    'HOME OWNER',
    'ADDRESS',
    'PACKAGES',
    { label: 'PRICE', align: 'right' },
    'APPROVAL STATUS',
    'CREATED'
  ];

  let store;
  let status = undefined;
  let variables;

  if (isAdmin) {
    headers.push('ACTIONS');
  }

  $: variables = {
    type: installationType,
    status,
    search: $page.url.searchParams.get('q'),
    betweenDates: $page.url.searchParams.has('dt')
      ? JSON.parse($page.url.searchParams.get('dt'))
      : thisMonth(),
    _v: 0
  };

  function approvalChanged({ detail }) {
    if (!store) return;

    store.update(s => {
      const i = s.find(i => i.id === detail.id);
      if (i) {
        i.approval = detail.approval;
      }
      return s;
    });
  }

  function dtChange({ detail }) {
    variables.betweenDates = detail;
    $page.url.searchParams.set('dt', JSON.stringify(detail));
  }
</script>

<Page {title}>
  <svelte:fragment slot='filters'>
    <div class='flex flex-col sm:flex-row gap-x-2 justify-start'>
      <DateRange value={variables.betweenDates} on:change={dtChange} />
      <DDLStatus bind:value={status} />
    </div>
  </svelte:fragment>

  <DataTable
    query={QryApprovedInstallations}
    dataProp='approvedInstallations'
    pageSize={20}
    headers={headers}
    {variables}
    searchable
    cbSetStore={(s) => {
      store = s;
    }}
  >
    <svelte:fragment slot='row' let:item let:headers>
      <DataColumn header={headers[0]}>{item.ownerName}</DataColumn>
      <DataColumn header={headers[1]}>{item.ownerAddress}</DataColumn>
      <DataColumn header={headers[2]}>{item.pkg}</DataColumn>
      <DataColumn header={headers[3]} align='right'>{usdFormat(item.price)}</DataColumn>
      <DataColumn header={headers[4]}>{item.approval}</DataColumn>
      <DataColumn header={headers[5]}>{localDate(item.approvalAt)}</DataColumn>
      {#if (isAdmin)}
        <DataColumn header={headers[6]}>
          <PendingJobActions
            {item}
            {basePath}
            on:approvalChange={approvalChanged}
          />
        </DataColumn>
      {/if}
    </svelte:fragment>
  </DataTable>
</Page>
