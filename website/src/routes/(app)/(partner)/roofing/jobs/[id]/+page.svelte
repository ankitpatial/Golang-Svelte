<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import JobDetailForUser from '$lib/components/job/JobDetailForUser.svelte';
  import Page from '$lib/components/Page.svelte';

  function back() {
    goto(PATH.ROOFING_PARTNER_JOBS + $page.url.search);
  }
</script>

<Page title='Job Detail'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} tooltip='Go Back To Dashboard' />
  </svelte:fragment>

  <JobDetailForUser id={$page.params.id} />
</Page>
