<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { getContextClient, queryStore } from '@urql/svelte';
  import { QryPartnerJobStats } from '$lib/gql/job';
  import { page } from '$app/stores';
  import { DataColumn, TableQueryStore } from '$lib/components/DataTable';
  import PartnerListCountCol from './PartnerListCountCol.svelte';
  import { goto } from '$app/navigation';
  import ToggleActive from '$lib/components/partner/ToggleActive.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import IconVideoCamera from '$lib/assets/svg/IconVideoCamera.svelte';
  import type { PartnerType } from '$lib/graph/graphql';
  import PartnerListItemActions from '$lib/components/partner/PartnerListItemActions.svelte';

  export let partnerType: PartnerType;
  export let baseURL;
  export let showMobilSettings = false;

  const client = getContextClient();
  const dataProp = 'partnerJobStats';
  const headers = [
    'Roofing Partner',
    'New',
    'CUSTOMER CONTACTED',
    'JOB DETAILS CONFIRMED',
    'PERMITTING',
    'SCHEDULED',
    'IN PROGRESS',
    'INSTALLED',
    'INVOICED',
    'TOTAL',
    'ACTIVE',
    'ACTIONS'
  ];

  let store, data, q, skip, take, busy = false,
    variables = {
      partnerType,
      search: $page.url.searchParams.get('q') || '',
      skip: Number($page.url.searchParams.get('skip')),
      take: Number($page.url.searchParams.get('take') || 20)
    };

  $:store = queryStore<any>({
    client,
    query: QryPartnerJobStats,
    variables
  });

  function handleChange() {
    variables = {
      partnerType,
      search: $page.url.searchParams.get('q') || '',
      skip: Number($page.url.searchParams.get('skip')),
      take: Number($page.url.searchParams.get('take') || 20)
    };
  }
</script>

<TableQueryStore
  headers={headers}
  {store}
  {dataProp}
  take={variables.take}
  showSearch
  on:change={handleChange}
>
  <svelte:fragment slot='row' let:item let:headers>
    <DataColumn header={headers[0]}>
      <div class='flex flex-row space-x-2'>
        {item.name}
        {#if 'Active' !== item.status && 'InActive' !== item.status}
          <div class='badge badge-outline ml-5'>{item.status}</div>
        {/if}
      </div>
    </DataColumn>
    <DataColumn header={headers[1]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='New'
        count={item.newCount}
        countFlagged='{item.newCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[2]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='CustomerContacted'
        count={item.contactedCount}
        countFlagged='{item.contactedCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[3]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='JobDetailsConfirmed'
        count={item.confirmedCount}
        countFlagged='{item.confirmedCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[4]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='Permitting'
        count={item.permittingCount}
        countFlagged='{item.permittingCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[5]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='Scheduled'
        count={item.scheduledCount}
        countFlagged='{item.scheduledCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[6]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='InProgress'
        count={item.inProgressCount}
        countFlagged='{item.inProgressCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[7]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='Installed'
        count={item.installedCount}
        countFlagged='{item.installedCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[8]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        status='Invoiced'
        count={item.invoicedCount}
        countFlagged='{item.invoicedCountFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[9]}>
      <PartnerListCountCol
        id={item.id}
        name={item.name}
        count={item.total}
        countFlagged='{item.totalFlagged}'
      />
    </DataColumn>
    <DataColumn header={headers[10]}>
      <ToggleActive {store} partnerID={item.id} name={item.name} status={item.status} />
    </DataColumn>
    <DataColumn header={headers[11]}>
      <PartnerListItemActions basePath={baseURL} {partnerType} id={item.id} name={item.name} />
    </DataColumn>
  </svelte:fragment>
</TableQueryStore>

