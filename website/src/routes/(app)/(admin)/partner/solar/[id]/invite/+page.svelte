<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import PartnerOnboardInvite from '$lib/components/partner/PartnerOnboardInvite.svelte';
  import Page from '$lib/components/Page.svelte';

  const returnURL = PATH.PARTNER_SOLAR;
</script>

<Page title='Invite A Solar Partner'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={() => goto(returnURL)} />
  </svelte:fragment>

  <PartnerOnboardInvite
    id={$page.params.id}
    partnerType='SOLAR'
    returnURL={returnURL}
  />
</Page>
