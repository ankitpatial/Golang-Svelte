<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  -->
<script lang='ts'>
  import DataLoader from '$lib/components/data/DataLoader';
  import { getContextClient } from '@urql/svelte';
  import DataIntersection from './DataIntersection.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import { createEventDispatcher } from 'svelte';

  export let dataQry;
  export let dataProp;
  export let variables = {};
  export let searchable = false;

  const dispatch = createEventDispatcher();
  const loader = new DataLoader(getContextClient(), dataQry, dataProp, 10);
  const loading = loader.loading;
  const store = loader.store;

  for (const [key, value] of Object.entries(variables)) {
    loader.setParam(key, value);
  }
  loader.setParam('search', '');
  loader.load();

  function search(e) {
    const val = e.detail.value;
    loader.setParam('search', val);
    loader.load(true);
    dispatch('search', val);
  }

</script>

<div class='w-full' class:blur-sm={$loading}>
  {#if (searchable)}
    <Input
      value={variables.search || ''}
      placeholder='search...'
      on:change={search}
    />
  {/if}
  <div class='flex flex-col h-48 overflow-y-auto'>
    {#each $store.items as item}
      <div class='flex flex-row gap-2'>
        <slot name='row' {item}></slot>
      </div>
    {/each}
    <DataIntersection {loader} />
  </div>
</div>
