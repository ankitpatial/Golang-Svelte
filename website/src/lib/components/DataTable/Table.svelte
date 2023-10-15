<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';
  import Pagination from '$lib/components/Pagination.svelte';
  import InputSearch from '$lib/components/form/InputSearch.svelte';
  import { createEventDispatcher } from 'svelte';
  import PlaceholderLine from '$lib/components/placeholder/PlaceholderLine.svelte';
  import { smallScreen } from '$lib/stores/responsive.js';

  const dispatch = createEventDispatcher();

  export let headers: Array<string> = [];
  export let items: Array<object> = [];
  export let searchable = false;
  export let totalCount: number | undefined = undefined;
  export let hasNextPage = false;
  export let hasPreviousPage = false;
  export let showPager = false;
  export let loading = false;


  function searchChange({ detail }) {
    dispatch('change', detail);
  }

  function next() {
    dispatch('change');
  }

  function previous() {
    dispatch('change');
  }

</script>


<div class='border rounded-md'>
  {#if searchable}
    <div class='top-0 sticky'>
      <InputSearch on:change={searchChange} />
    </div>
  {/if}

  <table class='table table table-xs table-zebra w-full '>
    {#if !$smallScreen && headers && headers.length > 0}
      <thead>
      <tr>
        {#each headers as th, idx (idx)}
          <th>{th}</th>
        {/each}
      </tr>
      </thead>
    {/if}
    <tbody>
    {#if loading}
      {#each Array(items?.length > 3 ? items.length : 3).fill(0) as dummy}
        <tr>
          <td colspan={headers.length}>
            <PlaceholderLine />
          </td>
        </tr>
      {/each}
    {:else }
      {#each items as item, idx}
        {#if $smallScreen}
          <tr class='bg-white even:bg-gray-50'>
            <td class='bg-transparent border-0'>
              <table class='w-full mt-4'>
                <slot name='row' {idx} {headers} {item} />
              </table>
            </td>
          </tr>
        {:else }
          <tr>
            <slot name='row' {idx} {headers} {item} />
          </tr>
        {/if}
      {/each}
    {/if}
    </tbody>

    {#if (showPager && !$smallScreen && items?.length > 0)}
      <tfoot>
      <tr>
        <th colspan={headers.length}>
          <div class='flex w-full justify-between place-items-baseline'>
            <div>
              <h4>
                <span class='font-normal'> Total Count:  </span>
                {totalCount}
              </h4>
            </div>
            <div>
              <Pagination
                hasNext={hasNextPage}
                hasPrevious={hasPreviousPage}
                on:next={next}
                on:previous={previous}
              />
            </div>
          </div>
        </th>
      </tr>
      </tfoot>
    {/if}
  </table>
  {#if $smallScreen && hasNextPage}
    <br />
    <div class='flex justify-between'>
      <IntersectionObserver on:intersect={next}>
        <div class='btn btn-outline' class:loading={next}>Loading...</div>
      </IntersectionObserver>
    </div>
  {/if}
</div>


<style>
  .tr-even tbody > tr:nth-child(even) th,
  .tr-even tbody > tr:nth-child(even) td {
    @apply bg-slate-50;
  }
</style>
