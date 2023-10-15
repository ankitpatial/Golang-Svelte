<script lang='ts'>
  import { fade } from 'svelte/transition';
  import DataIntersection from '$lib/components/data/DataIntersection.svelte';
  import DataLoader from '$lib/components/data/DataLoader';
  import { getContextClient } from '@urql/svelte';
  import { type InstallationType, ProductType } from '$lib/graph/graphql';
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import Package from './Package.svelte';
  import Book from '$lib/components/installation/Book.svelte';

  export let dataQry;
  export let dataProp;
  export let installationType: InstallationType;
  export let pendingJobsPath;

  const dl = new DataLoader(getContextClient(), dataQry, dataProp, 20);
  const store = dl.store;
  const loading = dl.loading;
  dl.setParam('category', ProductType.SmartHome);
  dl.load();

  let pkg = undefined;

  onMount(async () => {
    setTimeout(() => {
      goto('#pkg0');
    }, 300);
  });

  function select(e) {
    pkg = e.detail.pkg;
  }
</script>

{#if (pkg)}
  <div in:fade={{duration: 300}}>
    <Book
      {installationType}
      {pkg}
      on:saved={() => {
        goto(pendingJobsPath);
      }}
      on:cancel={() => {
        pkg = undefined;
      }}
    />
  </div>
{:else}
  <div in:fade={{duration: 300}}>
    <div class='carousel carousel-center max-w-2xl p-4 space-x-4'>
      {#each $store.items as pkg, i}
        <div id='pkg{i}' class='carousel-item'>
          <Package idx={i} data={pkg} on:select={select} />
        </div>
      {/each}
      <div class='carousel-item'>
        <DataIntersection loader={dl} />
      </div>
    </div>
    <div class='flex justify-start w-full py-2 gap-2 ml-4'>
      {#each $store.items as pkg, i}
        <a href='#pkg{i}' class='btn btn-xs'>{i + 1}</a>
      {/each}
    </div>
  </div>
{/if}
