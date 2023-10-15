<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { PATH } from '$lib/enum';
  import PartnerList from '$lib/components/partner/PartnerList.svelte';
  import { PartnerType } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';
</script>

<Page title='Roofing EPC'>
  <PartnerList partnerType={PartnerType.Epc} baseURL={PATH.PARTNER_EPC} />
</Page>
