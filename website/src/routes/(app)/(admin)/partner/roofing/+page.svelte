<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { PATH } from '$lib/enum.js';
  import PartnerOnboardActions from '$lib/components/partner/PartnerOnboardActions.svelte';
  import PartnerList from '$lib/components/partner/PartnerList.svelte';
  import { PartnerType } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';
</script>

<Page title='Roofing Partners'>
  <svelte:fragment slot='buttons'>
    <PartnerOnboardActions baseURL={PATH.PARTNER_ROOFING} />
  </svelte:fragment>
  <PartnerList partnerType={PartnerType.Roofing} baseURL={PATH.PARTNER_ROOFING} />
</Page>
