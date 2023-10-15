<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import Btn from "$lib/components/form/Btn.svelte";

  export let tooltip = '';
  export let cls = '';
  export let primary = false;
  export let square = false;
  export let busy = false;
  export let type: 'button' | 'submit' | 'reset' = 'button'
</script>

<Btn small {type} {primary} {square} {tooltip} {cls} {busy} on:click>
  <slot></slot>
</Btn>

