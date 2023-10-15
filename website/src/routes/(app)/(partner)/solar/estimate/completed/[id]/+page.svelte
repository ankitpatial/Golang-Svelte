<!--
- Copyright (c) Simsaw Software Pvt. Ltd. 2023.
- Author: Ankit Patial <ankit@simsaw.com>
-->

<script lang='ts'>
  import {page} from '$app/stores';
  import {goto} from '$app/navigation';
  import {PATH} from '$lib/enum.js';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import type Job from '$lib/models/Job';
  import Page from '$lib/components/Page.svelte';
  import EstimateDetail from "$lib/components/estimate/EstimateDetail.svelte";

  const title = 'Estimate Detail';
  // pull job info along with measurement orders
  let job: Job;

  $: returnTo = $page.url.searchParams.get('return');

  function back() {
    goto(PATH.SOLAR_PARTNER_EST_COMPLETED + $page.url.search);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='Completed Estimates'/>
  </svelte:fragment>

  <EstimateDetail id={$page.params.id} on:load={({detail}) => job = detail }/>
</Page>

