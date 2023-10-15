<script lang='ts'>
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { authUser } from '$lib/stores/auth';
  import PartnerUser from '$lib/components/partnerUser/PartnerUser.svelte';
  import { page } from '$app/stores';
  import Page from '$lib/components/Page.svelte';

  const title = 'Edit User';

  function back() {
    goto(PATH.PARTNER_VIEW_USERS + $page.url.search);
  }
</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='View Users' />
  </svelte:fragment>

  <PartnerUser isEdit partnerID={$authUser.partner?.id} dashboardPath={PATH.PARTNER_VIEW_USERS} />
</Page>
