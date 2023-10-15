<script lang='ts'>
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import TrainingKinds from '$lib/components/training/TrainingKinds.svelte';
  import { createEventDispatcher } from 'svelte';
  import Page from '$lib/components/Page.svelte';

  export let title: string;
  export let partnerID: string;

  const dispatch = createEventDispatcher();

  let name = $page.url.searchParams.get('name');

  function back() {
    $page.url.searchParams.delete('name');
    dispatch('back');
  }

</script>


<Page title='{title}: <strong>{name}</strong>'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>

  <TrainingKinds {partnerID} />
</Page>
