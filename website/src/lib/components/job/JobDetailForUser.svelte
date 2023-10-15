<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import {usdFormat} from '$lib/utils/currency.js';
  import type Job from '$lib/models/Job';
  import {getContextClient} from '@urql/svelte';
  import {QryMyJob} from '$lib/gql/partner';
  import alerts from '$lib/stores/alerts';
  import JobProgressLogs from '$lib/components/job/JobProgressLogs.svelte';
  import JobFiles from '$lib/components/job/JobFiles.svelte';
  import JobNotes from "$lib/components/job/JobNotes.svelte";
  import H3WithBorder from "$lib/components/H3WithBorder.svelte";
  import HomeOwner from './HomeOwner.svelte';
  import Material from './Material.svelte';

  export let id: string;

  // pull job info along with measurement orders
  const client = getContextClient();
  let job: Job, busy = false, isCurrentLowSlopeOnly = false;

  $: if (id) {
    busy = true;
    client.query(QryMyJob, {id})
      .toPromise()
      .then((res) => {
        busy = false;
        // check for error and alert
        if (res.error) {
          alerts.error('Job Info', res.error.message);
          return;
        }
        job = res.data.myJob;
      });
  }

</script>

<div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={busy}>
  <HomeOwner data={job?.homeOwner}>
    <InputReadonly
      label='Price'
      value={job?.workOrderPrice ? usdFormat(job?.workOrderPrice) : 'N/A'}
    />
  </HomeOwner>
  <Material data={job?.estimate} />
  <CardWithBorderTitle title='Documents'>
    <JobFiles {id}/>
  </CardWithBorderTitle>
  <div class='card w-full sm:w-11/12 '>
    <div class='card-body'>
      <H3WithBorder> Notes</H3WithBorder>
      <JobNotes {id} assignedPartnerID={job?.contractor?.id}/>
    </div>
  </div>

  <div class='card w-full sm:w-11/12 '>
    <div class='card-body'>
      <H3WithBorder> Progress Logs</H3WithBorder>
      <JobProgressLogs {id}/>
    </div>
  </div>
</div>
