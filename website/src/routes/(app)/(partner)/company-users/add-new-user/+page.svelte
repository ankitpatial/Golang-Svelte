<script lang='ts'>
  import PartnerUser from '$lib/components/partnerUser/PartnerUser.svelte';
  import { authUser } from '$lib/stores/auth';
  import { PATH } from '$lib/enum';
  import Page from '$lib/components/Page.svelte';

  const title = 'Add New User';

</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<Page {title}>
  <PartnerUser partnerID={$authUser.partner?.id} dashboardPath={PATH.PARTNER_VIEW_USERS} />
</Page>
