<script lang='ts'>
  import {getContextClient} from '@urql/svelte';
  import BtnIcon from "$lib/components/form/BtnIcon.svelte";
  import IconVideoCamera from "$lib/assets/svg/IconVideoCamera.svelte";
  import {goto} from "$app/navigation";
  import {QryTrainingVideoKinds} from "$lib/gql/training";
  import {onMount} from "svelte";
  import alerts from "$lib/stores/alerts";
  import {page} from "$app/stores";
  import type {DocumentNode} from "graphql/language";
  import {wsMessage} from "$lib/stores/socket";
  import {Channel} from "$lib/graph/graphql";

  export let partnerID = undefined;
  export let qry: DocumentNode = QryTrainingVideoKinds;
  export let dataProp: string = 'trainingVideoKinds';

  const client = getContextClient();

  let busy = false, trainingVideoKinds = [];

  onMount(() => {
    load();
    return wsMessage.subscribe((msg) => {
      if (!msg) return;
      if (Channel.TrainingVideo === msg.channel) {
        load()
      }
    })
  })

  function load() {
    busy = true;
    client.query(qry, {}).then((res) => {
      busy = false;
      if (res.error) {
        alerts.error('Get Training Kinds', res.error.message);
        return
      }

      trainingVideoKinds = res.data[dataProp] || []
    });
  }

  function viewVideos(k) {
    $page.url.searchParams.set('kn', k.name);
    if (partnerID) {
      $page.url.searchParams.set('partnerID', partnerID);
    }
    goto(`${$page.url.pathname}/${k.id}${$page.url.search}`)
  }

</script>

<div class="flex flex-col space-y-2" class:blur-sm={busy}>
  {#each trainingVideoKinds || [] as k}
    <div>
      <h2>{k.name}</h2>
      <BtnIcon on:click={() => viewVideos(k)}>
        <IconVideoCamera/>
      </BtnIcon>
    </div>
  {/each}

  {#if !busy && !trainingVideoKinds?.length}
    <p>
      No training video available at this moment
    </p>
  {/if}
</div>


