<script lang='ts'>
  import type Entity from '$lib/models/Entity';
  import {createEventDispatcher} from 'svelte';
  import Label from './Label.svelte';

  export let name: string;
  export let label: string = '';
  export let placeholder: string = '';
  export let value: any;
  export let options: [Entity];
  export let multiple = false;
  export let disabled = false;
  export let required = false;
  export let tooltip = '';
  export let w = ' ';
  export let m = 'mb-3';
  export let handleChange: (event: Event) => any = null;

  const dispatch = createEventDispatcher();
  let id: string = crypto.randomUUID();
  let changeTO;

  function change(e) {
    handleChange && handleChange(e);

    if (changeTO) clearTimeout(changeTO);

    changeTO = setTimeout(() => {
      dispatch('change', {name, value});
    }, 700);
  }
</script>

<div class='form-control {w} {m}'>
  <Label uppercase {id} content={label}>
    {#if $$slots.alt}
      <slot name='alt'></slot>
    {/if}
  </Label>
  <div>
    {#each options || [] as op}
      <label class='label cursor-pointer inline-flex justify-start mr-5'>
        <input
          {id}
          type='radio'
          class='radio radio-sm radio-secondary mr-2'
          name={name}
          value={op.id}
          bind:group={value}
          on:change={change}
          {disabled}
          {required}
        />
        <slot name="option" option={op}>
          <span class='label-text'>{op.name}</span>
        </slot>
      </label>
    {/each}
  </div>
</div>

