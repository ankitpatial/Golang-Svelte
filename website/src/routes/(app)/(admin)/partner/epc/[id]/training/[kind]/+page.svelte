<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import VideoCourses from '$lib/components/training/VideoCourses.svelte';
  import Page from '$lib/components/Page.svelte';

  const id = $page.params.id;
  const kind = $page.params.kind;
  const pageSize = 10;

  let title;
  let name;
  let kn;
  let partnerID;

  $: {
    kn = $page.url.searchParams.get('kn');
    name = $page.url.searchParams.get('name');
    title = kn + ' Videos : ' + name;
    partnerID = $page.url.searchParams.get('partnerID');
  }

  function back() {
    $page.url.searchParams.delete('partnerID');
    $page.url.searchParams.delete('kn');
    goto(`${PATH.PARTNER_EPC}/${id}/training${$page.url.search}`);
  }

</script>


<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>
  <VideoCourses {kind} {partnerID} allowAssign />
</Page>
