<script lang='ts'>
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';

  export let loader;
  const store = loader?.store;
  const loading = loader?.loading;
</script>

{#if !$loading && $store?.hasNext}
  <div class='flex justify-between'>
    <IntersectionObserver on:intersect={loader.loadMore}>
      <div class='text-xl'>
        <span class='-mt-2'>Loading</span>
        <span class='loading loading-dots loading-md '>&nbsp;</span>
      </div>
    </IntersectionObserver>
  </div>
{/if}
