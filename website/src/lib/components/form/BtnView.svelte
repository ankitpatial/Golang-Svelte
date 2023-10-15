<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import IconEyeOutline from '$lib/assets/svg/IconEyeOutline.svelte';
  import Btn from "$lib/components/form/Btn.svelte";

  export let tooltip = '';
  export let cls = '';
</script>

<Btn small ghost {tooltip} {cls} on:click>
  <IconEyeOutline/>
</Btn>
