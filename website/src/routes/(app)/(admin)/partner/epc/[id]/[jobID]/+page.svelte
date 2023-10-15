<script lang='ts'>
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import JobDetailForUser from '$lib/components/job/JobDetailForUser.svelte';
  import Page from '$lib/components/Page.svelte';

  const returnURL = `${PATH.PARTNER_EPC}/${$page.params.id}${$page.url.search}`;

  $: path = $page.url.pathname;
</script>

<Page title='Job Detail'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={() => goto(returnURL)} />
  </svelte:fragment>
  <JobDetailForUser id={$page.params.jobID} />
</Page>
