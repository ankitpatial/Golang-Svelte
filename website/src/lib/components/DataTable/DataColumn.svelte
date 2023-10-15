<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {smallScreen} from '$lib/stores/responsive.js';

  export let header = '';
  export let cls = '';
  export let clsMobile = 'w-32';
  export let whitespaceWrap = false;
  export let align: 'left' | 'right' | 'center' | 'justify' | 'char' = 'justify'
</script>

{#if $smallScreen}
  <tr>
    <td style="text-align: {align}">
      <div class='{clsMobile} whitespace-pre-line'>
        <label class='font-bold'>{header || ''}</label>
      </div>
    </td>
    <td class='whitespace-pre-wrap'>
      <slot></slot>
    </td>
  </tr>
{:else}
  <td class='{cls}' class:whitespace-pre-wrap={whitespaceWrap} style="text-align: {align}">
    <slot></slot>
  </td>
{/if}

