<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { goto } from '$app/navigation';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { PATH } from '$lib/enum';
  import PackageForm from '../PackageForm.svelte';
  import { onMount } from 'svelte';
  import Page from '$lib/components/Page.svelte';
  import { editPackage } from '../store';

  const title = 'Edit Product';
  const client = getContextClient();

  onMount(() => {
    if (!$editPackage) {
      back();
    }
  });

  function back() {
    goto(PATH.PRODUCT_PACKAGES);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Products' />
  </svelte:fragment>

  <div class='w-full sm:w-8/12 mt-10'>
    {#if ($editPackage)}
      <PackageForm on:saved={back} isEdit data={$editPackage} />
    {/if}
  </div>
</Page>

