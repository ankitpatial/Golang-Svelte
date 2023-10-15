<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {goto} from "$app/navigation";
  import {PATH} from "$lib/enum";
  import {page} from "$app/stores";
  import PartnerTrainingVideoKindsCourses from "$lib/components/training/PartnerTrainingVideoKindsCourses.svelte";

  const id = $page.params.id;

  function back() {
    goto(`${PATH.PARTNER_SOLAR}/${id}/training${$page.url.search}`)
  }
</script>

<PartnerTrainingVideoKindsCourses on:back={back}/>
