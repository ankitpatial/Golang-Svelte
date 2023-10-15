<script lang='ts'>
  import { QryAssignedJobs } from '$lib/gql/job';
  import { usdFormat } from '$lib/utils/currency';
  import { goto } from '$app/navigation';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import { page } from '$app/stores';
  import { DataColumn, DataTable } from '$lib/components/DataTable';
  import { thisMonth } from '$lib/utils/time';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import { localDate } from '$lib/utils/time.js';
  import Progress from '$lib/components/JobProgress';
  import { humanize } from '$lib/utils/string';
  import Modal from '$lib/components/Modal.svelte';
  import ProgressFilterDDL from '$lib/components/partner/ProgressFilterDDL.svelte';
  import { Channel, JobProgress, Topic } from '$lib/graph/graphql';
  import { onMount } from 'svelte';
  import { wsMessage } from '$lib/stores/socket';
  import Page from '$lib/components/Page.svelte';

  let count = 0,
    dt: Array<Date> = $page.url.searchParams.has('dt')
      ? JSON.parse($page.url.searchParams.get('dt'))
      : thisMonth(),
    progress = $page.url.searchParams.get('progress') || undefined,
    job,
    openModal = false,
    isProgressDirty = false;
  $:variables = {
    search: $page.url.searchParams.get('q'),
    progress,
    betweenDates: dt,
    orderBy: {
      direction: 'DESC',
      field: 'CREATED'
    },
    _v: 0
  };

  onMount(() => {
    return wsMessage.subscribe(msg => {
      if (!msg) return;
      if (Channel.Job === msg.channel && Topic.Progress === msg.topic) {
        variables._v = Date.now();
      }
    });
  });

  function gotoDetail(id) {
    goto(`${$page.url.pathname}/${id}` + $page.url.search);
  }

  function dtChange({ detail }) {
    dt = detail;
    $page.url.searchParams.set('dt', JSON.stringify(dt));
  }

  function viewProgress(item: any) {
    job = item;
    openModal = true;
  }

  function hideProgress() {
    job = undefined;
    openModal = false;
    if (isProgressDirty) {
      variables._v = Date.now();
    }
  }
</script>

<Page title='Assigned Jobs' subTitle='({count})'>
  <svelte:fragment slot='filters'>
    <div class='flex flex-col sm:flex-row gap-x-2 justify-start'>
      <DateRange value={dt} on:change={dtChange} />
      <ProgressFilterDDL bind:value={progress} />
    </div>
  </svelte:fragment>
  <div class='w-full sm:w-full'>
    <div class='card-body p-0'>
      <DataTable
        headers={[
        'HOME OWNER',
        'ADDRESS',
        'SOLAR COMPANY',
        'EPC',
        'INTEGRATION',
        'ROOFING PARTNER',
        'RFX PRICE',
        'ROOFING PARTNER PRICE',
        'DATE ASSIGNED',
        'STATUS',
        'JOB DETAILS'
      ]}
        query={QryAssignedJobs}
        dataProp='assignedJobs'
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
            {item.contractor?.name || ''}
          </DataColumn>
          <DataColumn header={headers[6]}>
            {usdFormat(item.price)}
          </DataColumn>
          <DataColumn header={headers[7]}>
            {usdFormat(item.workOrderPrice)}
          </DataColumn>
          <DataColumn header={headers[8]}>
            {localDate(item.createdAt, false)}
          </DataColumn>
          <DataColumn header={headers[9]}>
            <div class='link' on:click={()=> viewProgress(item)}>
              <div
                class:text-red-500={item.progressFlagged}
                class:text-yellow-500={item.progress === JobProgress.Delayed}
              >
                {humanize(item.progress || JobProgress.New).toUpperCase()}
              </div>
            </div>
          </DataColumn>
          <DataColumn>
            <BtnView on:click={()=> gotoDetail(item.id)} tooltip='View Detail' />
          </DataColumn>
        </svelte:fragment>
      </DataTable>
    </div>
  </div>

  <Modal
    hideOkBtn
    open={openModal}
    on:cancel={hideProgress}
    size='lg'
    title='{job?.ownerFirstName} {job?.ownerLastName}, {job?.streetNumber} {job?.streetName}, {job?.city}, {job?.state}'
    cancelText='CLOSE'
  >
    <Progress readonly {job} bind:isDirty={isProgressDirty} />
  </Modal>
</Page>
