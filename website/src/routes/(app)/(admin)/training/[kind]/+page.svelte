<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { getContextClient, queryStore } from '@urql/svelte';
  import alerts from '$lib/stores/alerts';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { QryNewULID } from '$lib/gql';
  import { page } from '$app/stores';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { QryTrainingVideoCourses } from '$lib/gql/training';
  import PlaceholderBasic from '$lib/components/placeholder/PlaceholderBasic.svelte';
  import VideoList from '$lib/components/training/VideoList.svelte';
  import Page from '$lib/components/Page.svelte';

  export let data;
  const client = getContextClient();
  const kind = $page.params.kind;
  const pageSize = 10;

  let name,
    title,
    store = queryStore({
      client,
      query: QryTrainingVideoCourses,
      variables: { kind, pageSize }
    });

  $: {
    name = $page.url.searchParams.get('kn');
    title = name + ' Videos';
  }

  function back() {
    goto(PATH.TRAINING_CENTER);
  }

  async function goToUpload() {
    const res = await client.query(QryNewULID).toPromise();
    if (res.error) {
      console.log('** Error', res.error.message);
      alerts.error('Upload Video', 'Please try again');
      return;
    }

    await goto(PATH.TRAINING_VIDEOS_SAVE.replace(':kind', $page.params.kind).replace(':id', res.data.newULID) + `?kn=${name}`);
  }

</script>

<Page {title}>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
    <button class='btn btn-outline btn-sm uppercase' on:click={goToUpload}>
      Upload a Video
    </button>
  </svelte:fragment>

  {#if $store.fetching}
    <PlaceholderBasic />
  {:else}
    <div class='flex flex-col space-y-2'>
      {#each $store.data?.trainingVideoCourses || [] as c}
        <h2>{c.name}</h2>
        <VideoList courseID={c.id} connection={c.videos} {kind} {pageSize} />
      {/each}
    </div>
  {/if}
</Page>
