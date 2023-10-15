<script lang='ts'>
  import { page } from '$app/stores';
  import InstallationJobDetail from '$lib/components/installation/InstallationJobDetail.svelte';
  import Page from '$lib/components/Page.svelte';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';

  const id = $page.params.id;

  function back() {
    goto(PATH.SMART_HOME_PENDING);
  }
</script>

<Page title='Pending Job Detail'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <InstallationJobDetail {id} />
</Page>
