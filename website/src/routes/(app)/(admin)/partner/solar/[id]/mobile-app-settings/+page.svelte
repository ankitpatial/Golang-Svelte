<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import MobileAppSettings from '$lib/components/partner/MobileAppSettings.svelte';
  import Page from '$lib/components/Page.svelte';

  let name,
    partnerID = $page.params.id,
    returnURL = PATH.PARTNER_SOLAR;

  $: path = $page.url.pathname;
  $: name = $page.url.searchParams.get('name');

  function back() {
    $page.url.searchParams.delete('step');
    goto(returnURL + $page.url.search);
  }
</script>

<Page title='{name} Mobile App Settings'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>

  <MobileAppSettings {partnerID} />
</Page>



