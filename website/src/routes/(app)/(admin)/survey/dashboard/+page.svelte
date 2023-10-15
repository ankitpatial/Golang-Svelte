<script lang='ts'>
  import { DataColumn, DataTable } from '$lib/components/DataTable';
  import { QrySurveys } from '$lib/gql/survey';
  import Page from '$lib/components/Page.svelte';

  let count = 0;
  $: variables = {
    search: '',
    pagination: { first: 20, after: null },
    orderBy: {
      direction: 'DESC',
      field: 'CREATED'
    }
  };

  function logItem(item: any) {
    console.log(item);
    return item;
  }
</script>

<Page title='Site Surveys' subTitle='({count})'>
  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <DataTable
        headers={[
    'DATE',
    'TIME SLOT',
    'NAME',
    'PHONE',
    'ADDRESS',
    'NOTES',
    'STATUS'
  ]}
        query={QrySurveys}
        dataProp='surveys'
        {variables}
        searchable
        bind:totalCount={count}
      >
        <svelte:fragment slot='row' let:item let:headers>
          <DataColumn header={headers[0]}>
            {item.date}
          </DataColumn>
          <DataColumn header={headers[1]}>
            {item.slot}
          </DataColumn>
          <DataColumn header={headers[2]}>
            {item.name}
          </DataColumn>
          <DataColumn header={headers[3]}>
            {item.phone}
          </DataColumn>
          <DataColumn header={headers[4]}>
            {item.address}
          </DataColumn>
          <DataColumn header={headers[5]}>
            {item.notes}
          </DataColumn>
          <DataColumn header={headers[6]}>
            {item.status}
          </DataColumn>
        </svelte:fragment>
      </DataTable>
    </div>
  </div>
</Page>
