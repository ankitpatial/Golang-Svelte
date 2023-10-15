<script lang='ts'>
  import { onMount } from 'svelte';

  export let ref = undefined;
  export let id: string = 'ta-' + Date.now();
  export let name = '';
  export let label = '';
  export let placeholder = '';
  export let required = false;
  export let value = '';
  export let disabled = false;
  export let focus = false;
  export let maxlength = 500;
  export let w = 'w-full';
  export let m = 'mb-3';
  export let err = '';

  export let handleChange: (event: Event) => any = null;
  const inputCls = 'textarea textarea-bordered w-full';

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;

  onMount(() => {
    if (focus) ref.focus();
  });

  function changeVal(e) {
    if (disabled) return;
    value = e.target.value;
    handleChange && handleChange(e);
  }
</script>

<div class='form-control {w} {m}'>
  <label class='label p-0 mb-0.5 uppercase' for={id}>
    <span>{label}</span>
    {#if $$slots.alt}
      <slot name='alt'></slot>
    {/if}
  </label>

  <div class='tooltip-top' class:tooltip={isTip} class:tooltip-error={isTipErr} data-tip={tip}>
		<textarea
      class='{inputCls}'
      class:textarea-error={isTipErr}
      bind:this={ref}
      id={id}
      {name}
      {required}
      {placeholder}
      {value}
      {disabled}
      on:input={changeVal}
      on:keypress {maxlength}></textarea>
  </div>
</div>

