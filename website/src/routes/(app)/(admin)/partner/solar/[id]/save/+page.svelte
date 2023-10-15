<script lang='ts'>
  import { afterNavigate, goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import OnboardNav from '$lib/components/onboard/solar/OnboardNav.svelte';
  import Onboard from '$lib/components/onboard/solar/Onboard.svelte';
  import Page from '$lib/components/Page.svelte';

  const returnURL = PATH.PARTNER_SOLAR;

  let step = 1;
  $: path = $page.url.pathname;

  afterNavigate((na) => {
    step = Number(na.to?.url.searchParams.get('step') || 1);
  });

</script>

<Page title='Save Solar Partner'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={() => goto(returnURL)} />
  </svelte:fragment>
  <svelte:fragment slot='titleContent'>
    <OnboardNav />
  </svelte:fragment>

  <Onboard {returnURL} {step} />
</Page>


