<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  10/08/22, 5:58 pm
  -->
<script lang='ts'>
  import InputSearch from '$lib/components/form/InputSearch.svelte';
  import Pagination from '$lib/components/Pagination.svelte';
  import DataLoader from '$lib/components/data/DataLoader';
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';
  import { smallScreen } from '$lib/stores/responsive.js';

  export let header = [];
  export let loader: DataLoader | undefined;
  export let searchable = false;
  export let loading = false;
  export let value = '';

  let store = loader?.store;
  let storeBusy = loader?.loading;
  let innerWidth = 1020;

</script>

<div class='border rounded-md'>
  <!-- table caption/ search -->
  {#if searchable}
    <div class='top-0 sticky'>
      <InputSearch bind:value />
    </div>
  {/if}

  <table class='table table-xs table-zebra w-full ' class:blur-sm={$storeBusy || loading}>
    {#if !$smallScreen && header && header.length > 0}
      <thead>
      <tr>
        {#each header as th, idx (idx)}
          <th>{th}</th>
        {/each}
      </tr>
      </thead>
    {/if}
    <tbody class='border-b border-b-2'>
    <slot>
      <tr>
        <td colspan={header.length}>
          <span>No Data Found</span>
        </td>
      </tr>
    </slot>
    </tbody>
    {#if (!$smallScreen && loader && $store.count > 0)}
      <tfoot>
      <tr>
        <th colspan={header.length}>
          <div class='flex w-full justify-between place-items-baseline'>
            <div>
              <h4>
                <span class='font-normal'> Total Count:  </span> {$store.count}
              </h4>
            </div>
            <div>
              <Pagination
                hasNext={$store.hasNext}
                hasPrevious={$store.hasPrevious}
                on:next={loader.nextPage}
                on:previous={loader.previousPage}
              />
            </div>
          </div>
        </th>
      </tr>
      </tfoot>
    {/if}
  </table>

  {#if $smallScreen && !$storeBusy && $store?.hasNext}
    <br />
    <div class='flex justify-between'>
      <IntersectionObserver on:intersect={loader?.loadMore}>
        <div class='btn btn-outline' class:loading={$storeBusy}>Loading...</div>
      </IntersectionObserver>
    </div>
  {/if}
</div>
