<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  09/05/22, 4:55 PM
  -->

<script lang='ts'>
  import IconCancel from '$lib/assets/svg/IconCancel.svelte';
  import { createEventDispatcher } from 'svelte';
  import Btn from '$lib/components/form/Btn.svelte';

  export let content;
  export let type: 'success' | 'info' | 'error' = 'info';
  export let hideCancel = false;
  export let cancelTooltip = '';
  export let cancelTooltipPosition: 'top' | 'right' | 'left' = 'top';

  let c, b, ct, ci;

  switch (type) {
    case 'info':
      c = 'slate';
      b = 'border-accent';
      ct = 'secondary';
      ci = 'currentColor';
      break;
    case 'success':
      c = 'green';
      b = 'border-green-300';
      ct = 'info';
      ci = 'green';
      break;
    case 'error':
      c = 'red';
      b = 'border-red-300';
      ct = 'error';
      ci = 'red';
      break;
  }

  const dispatch = createEventDispatcher();
  const style = `badge-${type || 'info'} hover:bg-${c}-200 bg-${c}-200 ${b} text-sm`;

  function cancel() {
    dispatch('cancel');
  }
</script>

<div
  class='badge  gap-2 mr-2 mt-1.5 pr-0 p-3 font-medium rounded-full subpixel-antialiased whitespace-nowrap {style}'
  class:pr-2={hideCancel}
>
  {content}
  {#if (!hideCancel)}
    <Btn circle ghost xs on:click={cancel} tooltip={cancelTooltip} tooltipPosition={cancelTooltipPosition}>
      <IconCancel className='h-4 w-4' color={ci} />
    </Btn>
  {/if}
</div>
