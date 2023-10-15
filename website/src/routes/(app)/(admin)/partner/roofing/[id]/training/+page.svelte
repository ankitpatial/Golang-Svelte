<script lang='ts'>
  import {goto} from '$app/navigation';
  import {PATH} from '$lib/enum';
  import {page} from '$app/stores';
  import PartnerTrainingVideoKinds from "$lib/components/training/PartnerTrainingVideoKinds.svelte";

  const title = 'Training Video Access';

  let partnerID = $page.params.id;

  function back() {
    goto(PATH.PARTNER_ROOFING + $page.url.search)
  }
</script>

<PartnerTrainingVideoKinds {title} {partnerID} on:back={back}/>
