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
  import PricingCalc from '$lib/components/PricingCalc.svelte';
  import Page from '$lib/components/Page.svelte';
  import EstimateDetail from "$lib/components/estimate/EstimateDetail.svelte";

  // pull job info along with measurement orders
  let job: Job;

  $: returnTo = $page.url.searchParams.get('return');

  function back() {
    goto(PATH.ESTIMATES_VIEW + $page.url.search);
  }
</script>

<Page title='Estimate Detail' subTitle={job?.status}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Estimates'/>
  </svelte:fragment>

  <EstimateDetail id={$page.params.id} on:load={({detail}) => job = detail }/>

  {#if Number($page.url.searchParams.get('calc')) === 1}
    <PricingCalc {job}/>
  {/if}
</Page>


