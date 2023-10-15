<!--
 Copyright (c) 2022. Ankit Patial.
 Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 Author: Ankit Patial
  -->
<script lang='ts'>
  import type Entity from '$lib/models/Entity';
  import Label from '$lib/components/form/Label.svelte';
  import {createEventDispatcher} from 'svelte';

  export let name = crypto.randomUUID();
  export let label = '';
  export let placeholder: string = '';
  export let value: any;
  export let options: [Entity];
  export let multiple = false;
  export let disabled = false;
  export let required = false;
  export let tooltip = '';
  export let cls = '';

  const dispatch = createEventDispatcher();

  let id: string = crypto.randomUUID();
  let changeTO;

  value = [].concat(value);

  function change() {
    if (changeTO) clearTimeout(changeTO);
    changeTO = setTimeout(() => {
      dispatch('change', {name, value});
    }, 700);
  }
</script>

<div class='form-control mb-3 {cls}'>
  <Label uppercase {id} content={label}>
    {#if $$slots.alt}
      <slot name='alt'></slot>
    {/if}
  </Label>
  <div class=" flex flex-col space-y-2 border rounded shadow-sm overflow-auto h-20">
    {#each options || [] as op}
      <label class='mt-2 ml-3 cursor-pointer inline-flex justify-start'>
        <input
          {name}
          type='checkbox'
          class='checkbox checkbox-sm checkbox-secondary mr-2'
          value={op.id}
          bind:group={value}
          on:change={change}
          {disabled}
        />
        <span class='label-text'>{op.name}</span>
      </label>
    {/each}
  </div>
</div>
