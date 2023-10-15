<script lang='ts'>
  import {createEventDispatcher} from 'svelte';
  import Label from './Label.svelte';

  export let label = '';
  export let name = '';
  export let required = false;
  export let disabled = false;
  export let value = undefined;
  export let tooltip = '';
  export let wrapperCls = ''
  export let inpCls = ''
  export let tooltipForLabel = ''

  const dispatch = createEventDispatcher();

  let inp;
  let id = crypto.randomUUID();

  $: if (inp && value == undefined) {
    inp.indeterminate = true;
  }

  function changeVal(e) {
    if (disabled) return;
    const v = e.target.checked;
    if (value !== v) {
      value = v;
      dispatch('change', v);
    }
  }
</script>

<div class="form-control {wrapperCls}">
  <div class="flex justify-between cursor-pointer">
    {#if label}
      <Label tooltip={tooltipForLabel} content={label} {id} uppercase/>
    {/if}
    <input
      {id}
      bind:this={inp}
      type='checkbox'
      class='toggle ml-2 tooltip-right {inpCls}'
      class:toggle-success={true}
      {name}
      {required}
      {disabled}
      checked={value}
      on:change={changeVal}
      on:keypress
      class:tooltip={!!tooltip} data-tip={tooltip}
    />
  </div>
</div>

