<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import IconChevronRight from '$lib/assets/svg/IconChevronRight.svelte';
  import IconChevronDown from '$lib/assets/svg/IconChevronDown.svelte';

  export let title = '';
  export let open = false;
  export let titleUnderline = true;

  function toggle() {
    open = !open;
  }
</script>

<div
  class='collapse collapse-plus inline-grid w-full'
  class:collapse-open={open}
  class:collapse-close={!open}>
  <div
    class='collapse-title text-sm text-accent mt-2 py-3 px-2 uppercase whitespace-nowrap flex hover:bg-secondary-focus/50 focus:text-primary rounded-md'
    class:border-b={titleUnderline}
    class:border-secondary-focus={titleUnderline}
    on:click|stopPropagation={toggle}
  >
    <span class='icon-plus'><IconChevronRight cls='h-5' /></span>
    <span class='icon-minus'><IconChevronDown cls='h-5' /></span>
    {title}
  </div>
  {#if (open)}
    <div class='collapse-content pr-0'>
      <slot></slot>
    </div>
  {/if}
</div>


<style>
  .collapse-plus .collapse-title:after {
    @apply text-accent -mt-3 text-lg;
    content: '';
  }

  .collapse-open.collapse-plus .collapse-title:before,
  .collapse-plus:focus:not(.collapse-close) .collapse-title:before,
  .collapse-plus:not(.collapse-close) input[type="checkbox"]:checked ~ .collapse-title:before {
    background-image: url("../../assets/svg/IconFolderPlus.svg");
  }

  .collapse-plus .collapse-title:before {
    @apply text-accent -mt-3 text-lg;
    /*content: '\203A';*/
  }


</style>
