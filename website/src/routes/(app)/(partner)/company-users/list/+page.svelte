<script lang='ts'>
  import { authUser } from '$lib/stores/auth';
  import PartnerUserList from '$lib/components/partnerUser/PartnerUserList.svelte';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';

  const title = 'View All Users';
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<Page {title}>
  <PartnerUserList partnerID={$authUser.partner?.id} editPath={PATH.PARTNER_EDIT_USERS} />
</Page>
