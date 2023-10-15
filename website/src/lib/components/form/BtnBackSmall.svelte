<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import IconArrowLeft from '$lib/assets/svg/IconArrowLeft.svelte';
  import BtnSmall from "$lib/components/form/BtnSmall.svelte";

  export let tooltip = '';
  export let cls = '';
</script>

<BtnSmall {tooltip} {cls} square on:click>
  <IconArrowLeft cls='h-5 w-5'/>
</BtnSmall>
