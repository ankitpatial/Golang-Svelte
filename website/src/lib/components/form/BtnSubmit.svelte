<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import Btn from "$lib/components/form/Btn.svelte";

  export let loading = false;
  export let cls = '';
</script>

<Btn type="submit" {cls} primary disabled={loading} busy={loading} on:click>
  <slot></slot>
</Btn>
