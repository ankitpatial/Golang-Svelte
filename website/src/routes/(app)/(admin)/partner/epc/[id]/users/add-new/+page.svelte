<script lang='ts'>
  import PartnerUser from '$lib/components/partnerUser/PartnerUser.svelte';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import Page from '$lib/components/Page.svelte';

  const title = `${$page.url.searchParams.get('name')} Add User`;
  const partnerID = $page.params.id;
  const dashboardPath = `${PATH.PARTNER_EPC}/${partnerID}/users${$page.url.search}`;

</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={()=> goto(dashboardPath)} />
  </svelte:fragment>
  <PartnerUser {partnerID} {dashboardPath} />
</Page>
