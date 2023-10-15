<script lang='ts'>
  import {createEventDispatcher} from 'svelte';

  export let label = '';
  export let labelPosition: 'right' | 'left' = 'left';
  export let name = '';
  export let required = false;
  export let disabled = false;
  export let value = undefined;
  export let tooltip = '';
  export let wrapperCls = ''
  export let inpCls = ''
  export let handleChange: (event: Event) => any = null;

  const dispatch = createEventDispatcher();

  let inp;
  let id = crypto.randomUUID();

  $: if (inp && value == undefined) {
    inp.indeterminate = true;
  }

  function changeVal(e) {
    if (disabled) return;

    handleChange && handleChange(e);

    const v = e.target.checked;
    if (value !== v) {
      value = v;
      dispatch('change', v);
    }
  }
</script>

<div class="form-control {wrapperCls}">
  <label class="label cursor-pointer">
    {#if label && labelPosition === 'left'}
      <span class="label-text">{label}</span>
    {/if}
    <input
      {id}
      bind:this={inp}
      type='checkbox'
      class='checkbox checkbox-primary ml-2 tooltip-right {inpCls}'
      class:toggle-success={true}
      {name}
      {required}
      {disabled}
      checked={value}
      on:change={changeVal}
      on:keypress
      class:tooltip={!!tooltip} data-tip={tooltip}
    />
    {#if label && labelPosition === 'right'}
      <span class="label-text">{label}</span>
    {/if}
  </label>
</div>

