<script lang='ts'>
  import { afterNavigate, goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import PartnerOnboardNav from '$lib/components/onboard/roofing/PartnerOnboardNav.svelte';
  import PartnerOnboard from '$lib/components/onboard/roofing/PartnerOnboard.svelte';
  import Page from '$lib/components/Page.svelte';

  const returnURL = PATH.PARTNER_ROOFING;

  let step = 1;
  $: path = $page.url.pathname;

  afterNavigate((na) => {
    step = Number(na.to?.url.searchParams.get('step') || 1);
  });

</script>

<Page title='Save Roofing Partner'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={() => goto(returnURL)} tooltip='Go Back' />
  </svelte:fragment>
  <svelte:fragment slot='titleContent'>
    <PartnerOnboardNav />
  </svelte:fragment>
  <PartnerOnboard {returnURL} {step} />
</Page>


