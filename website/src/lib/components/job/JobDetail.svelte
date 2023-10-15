<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { page } from '$app/stores';
  import InputReadonly from '$lib/components/form/InputReadonly.svelte';
  import { getContextClient } from '@urql/svelte';
  import { QryJob } from '$lib/gql/job';
  import alerts from '$lib/stores/alerts';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import { createEventDispatcher } from 'svelte';
  import { usdFormat } from '$lib/utils/currency.js';
  import Polygon from '$lib/components/maps/Polygon.svelte';
  import JobFiles from '$lib/components/job/JobFiles.svelte';
  import { type Job } from '$lib/graph/graphql';
  import JobNotes from '$lib/components/job/JobNotes.svelte';
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';
  import PriceSummary from '$lib/components/estimate/PriceSummary.svelte';
  import HomeOwner from './HomeOwner.svelte';
  import SalesRep from './SalesRep.svelte';
  import Material from '$lib/components/job/Material.svelte';

  export let id: string;

  // pull job info along with measurement orders
  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let job: Job;
  let busy = false;
  let boundary;

  $: if (id) {
    busy = true;
    client.query(QryJob, { id })
      .toPromise()
      .then((res) => {
        busy = false;
        // check for error and alert
        if (res.error) {
          alerts.error('Job Info', res.error.message);
          return;
        }
        job = res.data.job;
      });
  }

  $: dispatch('load', job);
  $: pathname = $page.url.pathname;
  $: isCurrentLowSlopeOnly = job?.estimate?.currentMaterial === 'Low Slope Only';

</script>

<div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={busy}>
  <HomeOwner data={job?.homeOwner}>
    <InputReadonly label='Price' value={usdFormat(job?.price)} />
    <InputReadonly label='Roofing Partner Price' value={usdFormat(job?.workOrderPrice)} />
  </HomeOwner>
  <SalesRep company={job?.companyName} data={job?.salesRep} />
  <Material data={job?.estimate} />
  <CardWithBorderTitle title='Documents'>
    <JobFiles {id} />
    <PriceSummary summary={job?.estimate?.priceSummary} />
  </CardWithBorderTitle>

  <div class='card w-full sm:w-11/12'>
    <div class='card-body'>
      <H3WithBorder> Notes</H3WithBorder>
      <JobNotes {id} assignedPartnerID={job?.partner?.id} />
    </div>
  </div>

</div>

{#if job?.latitude && job?.longitude}
  <div class='card w-full sm:w-11/12 mt-4'>
    <div class='card-body'>
      <H3WithBorder> Map</H3WithBorder>
      <Polygon position={{lat:job?.latitude, lng: job?.longitude }} readonly />
    </div>
  </div>
{/if}

<div class='clear-both mt-10 text-sm'>
  Requester: {job?.creatorName}
</div>
