<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import Btn from "$lib/components/form/Btn.svelte";

  export let tooltip = '';
  export let tooltipPosition: 'top' | 'right' | 'left' | 'bottom' = 'top'
  export let cls=''
</script>

<Btn square small ghost on:click {tooltip} {tooltipPosition} {cls}>
  <slot></slot>
</Btn>
