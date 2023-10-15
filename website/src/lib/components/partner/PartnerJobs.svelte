<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import {QryPartnerJobs} from '$lib/gql/job';
  import {page} from '$app/stores';
  import {DataColumn, DataTable} from '$lib/components/DataTable';
  import {usdFormat} from '$lib/utils/currency';
  import {localDate} from '$lib/utils/time';
  import BtnView from '$lib/components/form/BtnView.svelte';
  import {goto} from '$app/navigation';
  import {humanize} from '$lib/utils/string.js';
  import Progress from '$lib/components/JobProgress';
  import Modal from '$lib/components/Modal.svelte';
  import {Channel, JobProgress, Topic} from "$lib/graph/graphql";
  import {onMount} from "svelte";
  import {wsMessage} from "$lib/stores/socket";

  export let baseURL;
  export let partnerID = "";
  export let dates = [];
  export let flagged: boolean | undefined;
  export let progress: JobProgress | undefined;

  const client = getContextClient();
  const dataProp = 'partnerJobs';
  const searchProp = 'qpj';
  const headers = [
    'HOME OWNER',
    'STATUS',
    'EMAIL',
    'PHONE',
    'ADDRESS',
    'SOLAR COMPANY',
    'EPC',
    'ROOFING PARTNER',
    'PRICE',
    'INSTALL DATE',
    'INSPECTION DATE',
    ''
  ];

  let variables,
    job,
    isProgressDirty = false,
    openModal = false;

  $: variables = {
    partnerID,
    progress,
    flagged,
    dates,
    search: $page.url.searchParams.get(searchProp) || '',
    _v: 0
  }

  onMount(() => {
    return wsMessage.subscribe(msg => {
      if (!msg) return;
      if (Channel.Job === msg.channel && Topic.Progress === msg.topic) {
        variables._v = Date.now();
      }
    })
  })

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

<DataTable
  {headers}
  query={QryPartnerJobs}
  {dataProp}
  {searchProp}
  {variables}
  searchable
>
  <svelte:fragment slot='row' let:item let:headers>
    <DataColumn header={headers[0]}>
      {item.homeOwner?.firstName} {item.homeOwner?.lastName}
    </DataColumn>
    <DataColumn header={headers[1]}>
      <div class='link' on:click={()=> viewProgress(item)}>
        <div
          class:text-red-500={item.progressFlagged}
          class:text-yellow-500={item.progress === JobProgress.Delayed}
        >
          {humanize(item.progress || JobProgress.New).toUpperCase()}
        </div>
      </div>
    </DataColumn>
    <DataColumn header={headers[2]}>
      {item.salesRep?.email}
    </DataColumn>
    <DataColumn header={headers[3]}>
      {item.salesRep?.phone}
    </DataColumn>
    <DataColumn header={headers[4]}>
      {item.homeOwner?.streetNumber} {item.homeOwner?.streetName},<br>{item.homeOwner?.city}, {item.homeOwner?.state}
      - {item.homeOwner?.zip}
    </DataColumn>
    <DataColumn header={headers[5]}>
      {item.companyName}
    </DataColumn>
    <DataColumn header={headers[6]}>
      {item.epcName}
    </DataColumn>
    <DataColumn header={headers[7]}>
      {item.contractor?.name}
    </DataColumn>
    <DataColumn header={headers[8]}>
      {item.workOrderPrice > 0 ? usdFormat(item.workOrderPrice) : 'N/A'}
    </DataColumn>
    <DataColumn header={headers[9]}>
      {localDate(item.installDate, false)}
    </DataColumn>
    <DataColumn header={headers[10]}>
      {localDate(item.inspectionDate, false)}
    </DataColumn>
    <DataColumn>
      <BtnView on:click={() => goto(`${baseURL}/${item.id}${$page.url.search}`)}/>
    </DataColumn>
  </svelte:fragment>
</DataTable>

<Modal
  hideOkBtn
  open={openModal}
  on:cancel={hideProgress}
  size='lg'
  title='{job?.homeOwner?.firstName} {job?.homeOwner?.lastName}, {job?.homeOwner?.streetNumber} {job?.homeOwner?.streetName}, {job?.homeOwner?.city}, {job?.homeOwner?.state}'
  cancelText='CLOSE'
>
  <Progress {job} bind:isDirty={isProgressDirty}/>
</Modal>




