<script lang='ts'>
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import TrainingKinds from '$lib/components/training/TrainingKinds.svelte';
  import Page from '$lib/components/Page.svelte';

  const title = 'Training Video Access';

  let name = $page.url.searchParams.get('name');
  let partnerID = $page.params.id;

  function back() {
    $page.url.searchParams.delete('name');
    goto(PATH.PARTNER_INTEGRATION + $page.url.search);
  }
</script>

<Page label='{title}: <strong>{name}</strong>'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <TrainingKinds {partnerID} />
</Page>
