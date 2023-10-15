<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {QryUnassignedJobs} from '$lib/gql/job';
  import {usdFormat} from '$lib/utils/currency.js';
  import {goto} from '$app/navigation';
  import {DataColumn, DataTable} from '$lib/components/DataTable';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconUserPlus from '$lib/assets/svg/IconUserPlus.svelte';
  import {page} from '$app/stores';
  import {thisMonth} from '$lib/utils/time';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import {localDate} from '$lib/utils/time.js';
  import Page from '$lib/components/Page.svelte';

  let count = 0;
  let dt: Array<Date> = $page.url.searchParams.has('dt')
    ? JSON.parse($page.url.searchParams.get('dt'))
    : thisMonth();
  let partner: { id: '', name: '' } | undefined;
  $:variables = {
    search: $page.url.searchParams.get('q'),
    betweenDates: dt,
    orderBy: {
      direction: 'DESC',
      field: 'CREATED'
    }
  };

  function assign(id) {
    goto(`${$page.url.pathname}/${id}/assign-partner` + $page.url.search);
  }

  function dtChange({detail}) {
    dt = detail;
    $page.url.searchParams.set('dt', JSON.stringify(dt));
  }

  function gotoDetail(id) {
    goto(`${$page.url.pathname}/${id}` + $page.url.search);
  }
</script>

<Page title='Unassigned Jobs' subTitle='({count})'>
  <svelte:fragment slot='filters'>
    <DateRange value={dt} on:change={dtChange}/>
  </svelte:fragment>

  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <DataTable
        headers={['HOME OWNER', 'ADDRESS', 'SOLAR COMPANY', "EPC", 'INTEGRATION', 'RFX PRICE', 'ROOFING PARTNER PRICE', 'DATE APPROVED', 'Job Details', 'Assign Partner']}
        query={QryUnassignedJobs}
        dataProp='unassignedJobs'
        {variables}
        searchable
        bind:totalCount={count}
      >
        <svelte:fragment slot='row' let:item let:headers>
          <DataColumn header={headers[0]}>
            {item.homeOwner?.firstName} {item.homeOwner?.lastName}
          </DataColumn>
          <DataColumn header={headers[1]}>
            {item.homeOwner?.streetNumber} {item.homeOwner?.streetName}<br>
            {item.homeOwner?.city} {item.homeOwner?.state}, {item.homeOwner?.zip}
          </DataColumn>
          <DataColumn header={headers[2]}>
            {item.companyName}
          </DataColumn>
          <DataColumn header={headers[3]}>
            {item.epcName}
          </DataColumn>
          <DataColumn header={headers[4]}>
            N/A
          </DataColumn>
          <DataColumn header={headers[5]}>
            {usdFormat(item.price)}
          </DataColumn>
          <DataColumn header={headers[6]}>
            {usdFormat(item.workOrderPrice)}
          </DataColumn>
          <DataColumn header={headers[7]}>
            {localDate(item.createdAt, false)}
          </DataColumn>
          <DataColumn header={headers[8]}>
            <BtnView on:click={()=> gotoDetail(item.id)} tooltip='View Detail'/>
          </DataColumn>
          <DataColumn header={headers[9]}>
            <BtnIcon on:click={() => assign(item.id)} tooltip='Assign Partner'>
              <IconUserPlus/>
            </BtnIcon>
          </DataColumn>
        </svelte:fragment>
      </DataTable>
    </div>
  </div>
</Page>
