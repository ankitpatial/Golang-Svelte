<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {getContextClient, queryStore} from '@urql/svelte';
  import {QryMyTrainingVideoCourses, QryTrainingVideoCourses} from "$lib/gql/training";
  import PlaceholderBasic from "$lib/components/placeholder/PlaceholderBasic.svelte";
  import VideoList from "$lib/components/training/VideoList.svelte";
  import {onMount} from "svelte";
  import {wsMessage} from "$lib/stores/socket";
  import {Channel} from "$lib/graph/graphql";

  // props, required
  export let kind;
  // props, optional
  export let partnerID = undefined;
  export let allowAssign = false;
  export let myVideos = false

  const client = getContextClient();
  const pageSize = 10;

  const qry = myVideos ? QryMyTrainingVideoCourses : QryTrainingVideoCourses
  const dataProp = myVideos ? 'myTrainingVideoCourses' : 'trainingVideoCourses'

  let name,
    title,
    store = load();

  onMount(() => {
    return wsMessage.subscribe((msg) => {
      if (!msg) return;
      if (Channel.TrainingVideo === msg.channel) {
        if (kind === msg.data.kind) {
          store = load();
        }
      }
    })
  })


  function load() {
    return queryStore<any>({
      client,
      query: qry,
      variables: {kind, pageSize, partnerID},
    })
  }
</script>


{#if $store.fetching}
  <PlaceholderBasic/>
{:else}
  <div class="flex flex-col space-y-2">
    {#each $store.data[dataProp] || [] as c}
      <h2>{c.name}</h2>
      <VideoList courseID={c.id} connection={c.videos} {kind} {pageSize} {partnerID} {allowAssign} {myVideos}/>
    {/each}
  </div>
{/if}

