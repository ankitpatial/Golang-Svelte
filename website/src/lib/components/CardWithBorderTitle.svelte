<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script>
  import H3WithBorder from '$lib/components/H3WithBorder.svelte';

  export let title = '';
  export let busy = false;
  export let cls = 'sm:w-96';
</script>

<div class='card w-full {cls}' class:blur-sm={busy}>
  <div class='card-body flex-none'>
    {#if title}
      <H3WithBorder>
        {title}
        <slot name='btn'></slot>
      </H3WithBorder>
    {/if}
    <slot></slot>
  </div>
</div>
