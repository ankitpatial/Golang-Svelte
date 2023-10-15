<script lang='ts'>
  import PartnerUserList from '$lib/components/partnerUser/PartnerUserList.svelte';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import Page from '$lib/components/Page.svelte';

  const title = `${$page.url.searchParams.get('name')} Users`;
  const partnerID = $page.params.id;
  const editPath = `${PATH.PARTNER_INTEGRATION}/${partnerID}/users/edit`;

  function back() {
    $page.url.searchParams.delete('qu');
    goto(PATH.PARTNER_INTEGRATION + $page.url.search);
  }

  function addUser() {
    goto(`${PATH.PARTNER_INTEGRATION}/${partnerID}/users/add-new${$page.url.search}`);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
    <BtnPlusSmall on:click={addUser} />
  </svelte:fragment>
  <PartnerUserList {partnerID} {editPath} />
</Page>
