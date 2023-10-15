<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';
  import Pagination from '$lib/components/Pagination.svelte';
  import InputSearch from '$lib/components/form/InputSearch.svelte';
  import { page } from '$app/stores';
  import { writable } from 'svelte/store';
  import PlaceholderLine from '$lib/components/placeholder/PlaceholderLine.svelte';
  import { smallScreen } from '$lib/stores/responsive';
  import type { PageInfo, Pager } from './d';

  export let headers: Array<string> = [];
  export let query;
  export let dataProp = '';
  export let searchProp = 'q';
  export let variables = {};
  export let searchable = false;
  export let pageSize = 20;
  export let totalCount = 0;
  export let cbSetStore: (store) => void = () => {
  };

  let store = writable([]);
  const client = getContextClient();

  let pager: Pager = { first: pageSize },
    pageInfo: PageInfo,
    clear = true,
    setNodes = false,
    loading = false,
    err,
    searchInp;

  $: if (query) {
    load({ ...variables, page: { first: pageSize }, _v: undefined });
  }

  cbSetStore(store);

  function next() {
    clear = false;
    pager = { after: pageInfo.endCursor, first: pageSize };
    load({ ...variables, page: pager });
  }

  function previous() {
    clear = false;
    pager = { before: pageInfo.startCursor, last: pageSize };
    load({ ...variables, page: pager });
  }

  async function load(vars) {
    setNodes = false;
    loading = true;
    err = '';
    const res = await client.query(query, vars).toPromise();
    loading = false;
    if (res.error) {
      err = res.error.message;
      return;
    }

    const d = res.data[dataProp] || {};
    const newNodes = (d.edges || []).map((o: any) => o.node);
    pageInfo = d.pageInfo;
    totalCount = d.totalCount;

    if ($smallScreen) {
      store.set(clear ? newNodes : $store.concat(newNodes));
    } else {
      store.set(newNodes);
    }

    // make it true if in case false was
    clear = true;
  }

  function searchChange({ detail }) {
    const d = detail || '';
    if (d) {
      $page.url.searchParams.set(searchProp, d);
    } else {
      $page.url.searchParams.delete(searchProp);
    }

    // goto($page.url.search);
  }
</script>

<div class='border rounded-md'>
  <div class='top-0 sticky' class:hidden={!searchable}>
    <InputSearch bind:value={variables.search} on:change={searchChange} bind:ref={searchInp} />
  </div>
  <table class='table table table-xs table-zebra w-full'>
    {#if !$smallScreen && headers && headers.length > 0}
      <thead>
      <tr>
        {#each headers as th, idx (idx)}
          {#if typeof (th) === 'string'}
            <th>{th}</th>
          {:else if typeof (th) === 'object'}
            <th style='text-align: {th.align || "justify"}'>{th.label}</th>
          {/if}
        {/each}
      </tr>
      </thead>
    {/if}
    <tbody>
    {#if loading && clear}
      {#each Array($store?.length > 3 ? $store?.length : 3).fill(0) as dummy}
        <tr>
          <td colspan={headers.length}>
            <PlaceholderLine />
          </td>
        </tr>
      {/each}
    {:else }
      {#each $store || [] as item}
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
    {/if}
    </tbody>

    {#if (!$smallScreen && $store?.length > 0)}
      <tfoot>
      <tr>
        <th colspan={headers.length}>
          <div class='flex w-full justify-between place-items-baseline'>
            <div>
              <h4>
                <span class='font-normal'> Total Count:  </span> {totalCount}
              </h4>
            </div>
            <div>
              <Pagination
                hasNext={pageInfo?.hasNextPage}
                hasPrevious={pageInfo?.hasPreviousPage}
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
  {#if $smallScreen && pageInfo?.hasNextPage}
    <br />
    <div class='flex justify-between'>
      <IntersectionObserver on:intersect={next}>
        <div class='btn btn-outline' class:loading={loading}>Loading...</div>
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
