<script lang='ts'>
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import ProductForm from '../ProductForm.svelte';
  import { getContextClient, queryStore } from '@urql/svelte';
  import { QryNewULID } from '$lib/gql';
  import PlaceholderBasic from '$lib/components/placeholder/PlaceholderBasic.svelte';
  import Page from '$lib/components/Page.svelte';

  const title = 'Add New Product';
  const client = getContextClient();
  const store = queryStore({
    client,
    query: QryNewULID,
    variables: {}
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
    {#if ($store.fetching) }
      <PlaceholderBasic />
    {:else if ($store.error)}
      <p class='badge badge-error'>{$store.error}</p>
    {:else }
      <ProductForm id={$store.data.newULID} on:saved={back} />
    {/if}
  </div>
</Page>

