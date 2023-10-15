<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang="ts">
  import VideoItem from "./VideoItem.svelte";
  import {getContextClient} from "@urql/svelte";
  import {QryMyTrainingVideos, QryTrainingVideos} from "$lib/gql/training";
  import alerts from "$lib/stores/alerts";
  import IntersectionObserver from "$lib/components/IntersectionObserver.svelte";
  import PlaceholderTextSmall from "$lib/components/placeholder/PlaceholderTextSmall.svelte";

  // props, required
  export let kind: string;
  export let courseID: string;
  export let pageSize: number;
  export let connection;
  // props, optional
  export let partnerID = '';
  export let allowAssign = false;
  export let myVideos = false

  const client = getContextClient();
  const qry = myVideos ? QryMyTrainingVideos : QryTrainingVideos
  const dataProp = myVideos ? 'myTrainingVideos' : 'trainingVideos'

  let busy = false,
    pageInfo = connection.pageInfo,
    edges = connection.edges;

  async function next() {
    if (busy || !pageInfo?.hasNextPage) return

    const page = {after: pageInfo.endCursor, first: pageSize}
    busy = true;
    const res = await client.query(qry, {kind, courseID, partnerID, page}).toPromise();
    busy = false
    if (res.error) {
      alerts.error('Error', res.error.message);
      return;
    }
    const d = res.data[dataProp] || {};
    pageInfo = d.pageInfo;
    edges = edges.concat(d.edges);
  }
</script>

<div class="overflow-hidden">
  <div class="carousel carousel-center max-w-lg p-4 space-x-4 bg-gray-200">
    {#each edges || [] as e, key(e.node.id)}
      <div class="carousel-item">
        <VideoItem video={e.node} {partnerID} {allowAssign}/>
      </div>
    {/each}
    {#if (!busy && pageInfo?.hasNextPage)}
      <div class="carousel-item">
        <IntersectionObserver on:intersect={next}>
          <PlaceholderTextSmall/>
        </IntersectionObserver>
      </div>
    {/if}
  </div>

</div>
