<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import JobDetail from '$lib/components/job/JobDetail.svelte';
  import Page from '$lib/components/Page.svelte';
  import type { Job } from '$lib/graph/graphql';

  let job: Job;

  function back() {
    goto(PATH.JOBS_UNASSIGNED + $page.url.search);
  }
</script>


<Page title='Unassigned Job Detail' subTitle={job?.status}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='Back to Unassigned Jobs' />
  </svelte:fragment>

  <JobDetail id={$page.params.id} on:load={({detail}) => job = detail } />
</Page>

