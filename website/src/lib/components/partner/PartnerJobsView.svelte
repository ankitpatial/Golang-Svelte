<script lang='ts'>
  import { page } from '$app/stores';
  import { thisYear } from '$lib/utils/time';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import InputToggle from '$lib/components/form/InputToggle.svelte';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import PartnerJobs from './PartnerJobs.svelte';
  import ProgressFilterDDL from './ProgressFilterDDL.svelte';
  import { createEventDispatcher } from 'svelte';
  import SolarPartnerJobs from '$lib/components/partner/SolarPartnerJobs.svelte';
  import Page from '$lib/components/Page.svelte';

  export let title: string;
  export let partnerID: string = '';
  export let baseURL: string;
  export let hideBackBtn = false;
  export let isSolarPartner = false;

  const dispatch = createEventDispatcher();

  let
    progress = $page.url.searchParams.get('progress') || undefined,
    flagged = $page.url.searchParams.has('flagged')
      ? Boolean($page.url.searchParams.get('flagged'))
      : false,
    name = $page.url.searchParams.get('name') || 'Partner',
    dates: Array<Date> = $page.url.searchParams.has('dt')
      ? JSON.parse($page.url.searchParams.get('dt'))
      : thisYear(),
    show = true;

  function back() {
    show = false;
    $page.url.searchParams.delete('dt');
    $page.url.searchParams.delete('qpj');
    $page.url.searchParams.delete('name');
    $page.url.searchParams.delete('flagged');
    $page.url.searchParams.delete('progress');
    dispatch('back');
  }

  function dtChange({ detail }) {
    dates = detail;
    $page.url.searchParams.set('dt', JSON.stringify(dates));
  }
</script>


<Page {title}>
  <svelte:fragment slot='buttons'>
    {#if (!hideBackBtn)}
      <BtnBackSmall on:click={back} />
    {/if}
    <slot></slot>
  </svelte:fragment>
  <svelte:fragment slot='filters'>
    <div class='flex flex-col sm:flex-row gap-x-2 justify-start'>
      <InputToggle name='flagged' label='Flagged' wrapperCls='mt-2' inpCls='toggle-error' bind:value={flagged} />
      <DateRange value={dates} on:change={dtChange} />
      <ProgressFilterDDL bind:value={progress} skipPaymentOptions />
    </div>
  </svelte:fragment>

  {#if isSolarPartner}
    <SolarPartnerJobs {partnerID} {baseURL} {progress} {flagged} {dates} />
  {:else}
    <PartnerJobs {partnerID} {baseURL} {progress} {flagged} {dates} />
  {/if}
</Page>
