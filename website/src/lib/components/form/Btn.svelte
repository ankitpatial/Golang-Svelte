<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  export let type: 'button' | 'submit' | 'reset' = 'button';
  export let primary = false;
  export let ghost = false;
  export let circle = false;
  export let square = false;
  export let small = false;
  export let busy = false;
  export let xs = false;
  export let shadow = false;
  export let cls = '';
  export let transparent: boolean = false;
  export let tooltip = '';
  export let tooltipPosition: 'top' | 'right' | 'left' = 'top';
  export let disabled = false;
  const trans = transparent ? 'bg-transparent/20' : '';
</script>

<div class='tooltip-{tooltipPosition} {cls}' class:tooltip={!!tooltip} data-tip={tooltip}>
  <button
    {type}
    class='btn uppercase {cls} {trans} whitespace-normal'
    class:btn-primary={primary}
    class:btn-outline={!primary}
    class:btn-circle={circle}
    class:btn-square={square}
    class:btn-ghost={ghost}
    class:btn-sm={small}
    class:btn-xs={xs}
    class:btn-accent={transparent}
    class:disabled={disabled}
    class:shadow-md={shadow}
    on:click
    tabindex="0"
  >
    <slot></slot>
    {#if busy}
      <span class='loading loading-xs loading-bars'></span>
    {/if}
  </button>
</div>
