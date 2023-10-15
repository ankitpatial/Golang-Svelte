<script lang='ts'>
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import PartnerUser from '$lib/components/partnerUser/PartnerUser.svelte';
  import { page } from '$app/stores';
  import Page from '$lib/components/Page.svelte';

  const title = `${$page.url.searchParams.get('name')} Edit User`;
  const partnerID = $page.params.id;
  const dashboardPath = `${PATH.PARTNER_ROOFING}/${partnerID}/users${$page.url.search}`;

  function back() {
    goto(dashboardPath);
  }
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Users' />
  </svelte:fragment>
  <PartnerUser isEdit {partnerID} {dashboardPath} />
</Page>
