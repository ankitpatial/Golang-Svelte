<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import { goto } from '$app/navigation';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { PATH } from '$lib/enum';
  import ProductForm from '../ProductForm.svelte';
  import { editProduct } from '../store';
  import { onMount } from 'svelte';
  import Page from '$lib/components/Page.svelte';

  const title = 'Edit Product';
  const client = getContextClient();

  onMount(() => {
    if (!$editProduct) {
      back();
    }
  });

  function back() {
    goto(PATH.PRODUCTS);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Products' />
  </svelte:fragment>

  <div class='w-full sm:w-8/12 mt-10'>
    {#if ($editProduct)}
      <ProductForm on:saved={back} isEdit data={$editProduct} />
    {/if}
  </div>
</Page>

