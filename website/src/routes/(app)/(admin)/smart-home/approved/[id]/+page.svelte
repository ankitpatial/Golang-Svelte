<script lang='ts'>
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import InstallationJobDetail from '$lib/components/installation/InstallationJobDetail.svelte';

  const id = $page.params.id;

  function back() {
    goto(PATH.SMART_HOME_APPROVED);
  }
</script>

<Page title='Approved Job Detail'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <InstallationJobDetail {id} />
</Page>
