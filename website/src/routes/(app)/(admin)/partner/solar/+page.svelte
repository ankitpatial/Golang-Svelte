<script lang='ts'>
  import { PATH } from '$lib/enum.js';
  import PartnerOnboardActions from '$lib/components/partner/PartnerOnboardActions.svelte';
  import PartnerList from '$lib/components/partner/PartnerList.svelte';
  import { PartnerType } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';
</script>

<svelte:head>
  <title>Solar Partners</title>
</svelte:head>

<Page title='Solar Partners'>
  <svelte:fragment slot='buttons'>
    <PartnerOnboardActions baseURL={PATH.PARTNER_SOLAR} />
  </svelte:fragment>

  <PartnerList partnerType={PartnerType.Solar} baseURL={PATH.PARTNER_SOLAR} />
</Page>
