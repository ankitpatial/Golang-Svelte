<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import IconPlusSmall from '$lib/assets/svg/IconPlusSmall.svelte';
  import BtnSmall from "$lib/components/form/BtnSmall.svelte";

  export let tooltip = '';
</script>

<BtnSmall square on:click {tooltip}>
  <IconPlusSmall/>
</BtnSmall>
