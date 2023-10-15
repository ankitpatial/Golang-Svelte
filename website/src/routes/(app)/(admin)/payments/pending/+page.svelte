<script lang='ts'>
  import { QryPendingInvoiceJobs } from '$lib/gql/job';
  import { usdFormat } from '$lib/utils/currency.js';
  import { goto } from '$app/navigation';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import { page } from '$app/stores';
  import { DataColumn, DataTable } from '$lib/components/DataTable';
  import { endOfDay, thisMonth } from '$lib/utils/time';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import { localDate } from '$lib/utils/time.js';
  import JobProgressName from '$lib/components/job/JobProgressName.svelte';
  import Page from '$lib/components/Page.svelte';

  let count = 0,
    dateRange: Array<any> = thisMonth();
  $:variables = {
    search: $page.url.searchParams.get('q'),
    betweenDates: dateRange && dateRange.length == 2
      ? [dateRange[0].toISOString(), endOfDay(dateRange[1]).toISOString()]
      : [],
    orderBy: {
      direction: 'DESC',
      field: 'CREATED'
    }
  };

  function gotoDetail(id) {
    goto(`${$page.url.pathname}/${id}` + $page.url.search);
  }
</script>

<Page title='Payments, Pending Approval' subTitle='({count})'>
  <svelte:fragment slot='filters'>
    <DateRange bind:value={dateRange} />
  </svelte:fragment>

  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <DataTable
        headers={[
        'HOME OWNER',
        'STATUS',
        'EMAIL',
        'PHONE',
        'ADDRESS',
        'SOLAR COMPANY',
        'ROOFING PARTNER',
        'PRICE',
        'INSTALL DATE',
        'INSPECTION DATE',
        ''
        ]}
        query={QryPendingInvoiceJobs}
        dataProp='jobsByProgress'
        {variables}
        searchable
        bind:totalCount={count}
      >
        <svelte:fragment slot='row' let:item let:headers>
          <DataColumn header={headers[0]}>
            {item.ownerFirstName} {item.ownerLastName}
          </DataColumn>
          <DataColumn header={headers[1]}>
            <JobProgressName progress={item.progress} progressAt={item.progressAt} />
          </DataColumn>
          <DataColumn header={headers[2]}>
            {item.repEmail}
          </DataColumn>
          <DataColumn header={headers[3]}>
            {item.repMobile}
          </DataColumn>
          <DataColumn header={headers[4]}>
            {item.streetNumber} {item.streetName}<br>
            {item.city} {item.state}, {item.zip}
          </DataColumn>
          <DataColumn header={headers[5]}>
            {item.companyName}
          </DataColumn>
          <DataColumn header={headers[6]}>
            {item.partner?.name}
          </DataColumn>
          <DataColumn header={headers[7]}>
            {usdFormat(item.price)}
          </DataColumn>
          <DataColumn header={headers[8]}>
            {localDate(item.installDate, false)}
          </DataColumn>
          <DataColumn header={headers[9]}>
            {localDate(item.inspectionDate, false)}
          </DataColumn>
          <DataColumn>
            <BtnView on:click={()=> gotoDetail(item.id)} tooltip='View Detail' />
          </DataColumn>
        </svelte:fragment>
      </DataTable>
    </div>
  </div>
</Page>
