<script lang='ts'>
  import DataTable from '$lib/components/DataTable/DataTable.svelte';
  import { QryJobProgress } from '$lib/gql/job.progress';
  import DataColumn from '$lib/components/DataTable/DataColumn.svelte';
  import { localDate } from '$lib/utils/time.js';

  export let id: string;
</script>


<DataTable
  query={QryJobProgress}
  dataProp='jobProgress'
  variables={{id,search: ''}}
  headers={['Status', 'Date', 'Note']}
>
  <svelte:fragment slot='row' let:item let:headers>
    <DataColumn header={headers[0]}>
      {item.status}
    </DataColumn>
    <DataColumn header={headers[1]}>
      {localDate(item.statusAt)}
    </DataColumn>
    <DataColumn header={headers[2]}>
      {item.note}
    </DataColumn>
  </svelte:fragment>
</DataTable>
