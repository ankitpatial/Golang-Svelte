<script lang='ts'>
  import { afterNavigate, goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import OnboardNav from '$lib/components/onboard/solar/OnboardNav.svelte';
  import Onboard from '$lib/components/onboard/solar/Onboard.svelte';
  import Page from '$lib/components/Page.svelte';

  const name = $page.url.searchParams.get('name') + ' Profile';

  let step = 1;
  let returnURL = '';
  $: path = $page.url.pathname;
  $: returnURL = `${PATH.PARTNER_SOLAR}`;

  afterNavigate((na) => {
    step = Number(na.to?.url.searchParams.get('step') || 1);
  });

  function back() {
    $page.url.searchParams.delete('step');
    goto(returnURL + $page.url.search);
  }
</script>

<Page title={name}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <svelte:fragment slot='titleContent'>
    <OnboardNav />
  </svelte:fragment>

  <Onboard {returnURL} {step} />
</Page>


