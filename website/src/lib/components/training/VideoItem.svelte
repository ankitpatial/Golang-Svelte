<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang="ts">
  import Video from "$lib/components/Video.svelte";
  import InputToggle from "$lib/components/form/InputToggle.svelte";
  import {getContextClient} from "@urql/svelte";
  import {PartnerTrainingVideoAccess} from "$lib/gql/training";
  import alerts from "$lib/stores/alerts";

  // props, required
  export let video;
  // props, optional
  export let partnerID = '';
  export let allowAssign = false;

  // var
  const client = getContextClient()
  const poster = video?.posterURL;
  const src = video?.videoURL;

  // handle assign change
  async function handleChange({detail}) {
    const title = 'Assign Video'
    const res = await client.mutation(PartnerTrainingVideoAccess, {partnerID, videoID: video.id, enabled: detail}).toPromise();
    if (res.error) {
      alerts.error(title, res.error.message);
      return
    }

    alerts.success(title, "Successfully assigned training video");
  }
</script>

<div class="card card-compact w-56 bg-base-100 shadow-xl">
  <figure>
    <Video {poster} {src}/>
  </figure>
  <div class="card-body">
    <h2 class="card-title">{video?.title}</h2>
    <p>{video?.description}</p>
    {#if (allowAssign)}
      <div class="card-actions justify-end">
        <InputToggle value={video.assigned || false} label="Assign" on:change={handleChange}/>
      </div>
    {/if}
  </div>
</div>
