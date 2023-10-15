<!--
- Copyright (c) Simsaw Software Pvt. Ltd. 2023.
- Author: Ankit Patial <ankit@simsaw.com>
-->

<script lang='ts'>
  import { page } from '$app/stores';
  import { thisMonth } from '$lib/utils/time';
  import DateRange from '$lib/components/form/dtp/DateRange.svelte';
  import EstStatusDDL from './EstStatusDDL.svelte';
  import EstimateList from './EstimateList.svelte';
  import Page from '$lib/components/Page.svelte';

  export let title;
  export let basePath;
  export let ignore = [];

  let store, totalCount = 0,
    dt: Array<Date> = $page.url.searchParams.has('dt')
      ? JSON.parse($page.url.searchParams.get('dt'))
      : thisMonth(),
    status = $page.url.searchParams.get('status') || '';

  function dtChange({ detail }) {
    dt = detail;
    $page.url.searchParams.set('dt', JSON.stringify(dt));
  }

  function statusChange({ detail }) {
    status = detail;
    $page.url.searchParams.set('status', status);
  }
</script>

<Page {title} subTitle='({totalCount})'>
  <svelte:fragment slot='filters'>
    <div class='flex flex-col sm:flex-row gap-x-2 justify-start'>
      <DateRange value={dt} on:change={dtChange} />
      <EstStatusDDL value={status} on:change={statusChange} {ignore} />
    </div>
  </svelte:fragment>

  <EstimateList {dt} {status} {ignore} bind:totalCount {basePath} />
</Page>

