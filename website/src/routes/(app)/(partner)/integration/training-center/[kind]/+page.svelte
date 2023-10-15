<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {goto} from "$app/navigation";
  import {PATH} from "$lib/enum";
  import MyTrainingVideoKindCourses from "$lib/components/training/MyTrainingVideoKindCourses.svelte";

  function back() {
    goto(PATH.ROOFING_PARTNER_TRAINING)
  }
</script>

<MyTrainingVideoKindCourses on:back={back}/>
