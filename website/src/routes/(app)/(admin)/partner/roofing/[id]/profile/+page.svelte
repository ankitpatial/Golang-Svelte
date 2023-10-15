<script lang='ts'>
  import { afterNavigate, goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import PartnerOnboardNav from '$lib/components/onboard/roofing/PartnerOnboardNav.svelte';
  import PartnerOnboard from '$lib/components/onboard/roofing/PartnerOnboard.svelte';
  import Page from '$lib/components/Page.svelte';

  const name = $page.url.searchParams.get('name') + ' Profile';
  const returnURL = `${PATH.PARTNER_ROOFING}/${$page.params.id}`;

  let step = 1;

  afterNavigate((na) => {
    step = Number(na.to?.url.searchParams.get('step') || 1);
  });

  function back() {
    $page.url.searchParams.delete('step');
    goto(PATH.PARTNER_ROOFING + $page.url.search);
  }
</script>

<Page title={name}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <svelte:fragment slot='titleContent'>
    <PartnerOnboardNav />
  </svelte:fragment>
  <PartnerOnboard {returnURL} {step} />
</Page>


