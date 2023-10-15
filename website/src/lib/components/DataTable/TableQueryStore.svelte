<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  10/08/22, 5:58 pm
  -->
<script lang='ts'>
  import { createEventDispatcher } from 'svelte';
  import InputSearch from '$lib/components/form/InputSearch.svelte';
  import Pagination from '$lib/components/Pagination.svelte';
  import { page } from '$app/stores';
  import { smallScreen } from '$lib/stores/responsive.js';

  export let headers = [];
  export let store;
  export let dataProp: string;
  export let take: number = 0;
  export let showSearch = false;
  export let hidePaging = false;

  const dispatch = createEventDispatcher();
  let items, hasNext, hasPrevious, skip, q;

  $:{
    q = $page.url.searchParams.get('q') || '';
    skip = Number($page.url.searchParams.get('skip'));
  }
  $: {
    items = ($store?.data && $store.data[dataProp]) || [];
    hasNext = take > 0 && items.length == take;
    hasPrevious = skip > 0;
  }

  function search({ detail }) {
    if (detail) {
      $page.url.searchParams.set('q', detail);
    } else {
      $page.url.searchParams.delete('q');
    }
    $page.url.searchParams.delete('skip');

    dispatch('change', $page.url.searchParams);
  }

  function previousPage() {
    if (!$store?.operation?.variables) {
      console.warn('$store.operation.variables is undefined');
      return;
    }

    skip -= take;
    if (skip < take) {
      skip = 0;
    }

    $page.url.searchParams.set('skip', skip);
    dispatch('change', $page.url.searchParams);
  }

  function nextPage() {
    if (!$store?.operation?.variables) {
      console.warn('$store.operation.variables is undefined');
      return;
    }

    if (skip === 0) {
      skip = take;
    } else {
      skip += take;
    }

    $page.url.searchParams.set('skip', skip);
    dispatch('change', $page.url.searchParams);
  }

</script>


<div class='border rounded-md w-full'>
  <!-- table caption/ search -->
  {#if (showSearch)}
    <div class='top-0 sticky'>
      <InputSearch value={q} on:change={search} />
    </div>
  {/if}

  <table class='table table-xs table-zebra w-full overflow-auto' class:blur-sm={$store?.fetching}>
    {#if !$smallScreen && headers && headers.length > 0}
      <thead>
      <tr>
        {#each headers as th, idx (idx)}
          <th>{th}</th>
        {/each}
      </tr>
      </thead>
    {/if}
    <tbody class='border-b border-b-2'>
    {#each items || [] as item}
      {#if $smallScreen}
        <tr class='bg-white even:bg-gray-50'>
          <td class='bg-transparent border-0'>
            <table class='w-full mt-4'>
              <slot name='row' {headers} {item} />
            </table>
          </td>
        </tr>
      {:else }
        <tr>
          <slot name='row' {headers} {item} />
        </tr>
      {/if}
    {/each}
    </tbody>
    {#if !hidePaging && (items.length > 0)}
      <tfoot>
      <tr>
        <th colspan={headers.length}>
          <div class='flex w-full justify-between place-items-baseline'>
            <div>
              <h4>

              </h4>
            </div>
            <div>
              <Pagination
                hasNext={hasNext}
                hasPrevious={hasPrevious}
                on:next={nextPage}
                on:previous={previousPage}
              />
            </div>
          </div>
        </th>
      </tr>

      </tfoot>
    {/if}
  </table>
</div>
