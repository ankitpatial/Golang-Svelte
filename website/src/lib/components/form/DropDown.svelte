<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  18/04/22, 2:52 PM
  -->
<script lang='ts'>
  import type Entity from '$lib/models/Entity';
  import {createEventDispatcher} from 'svelte';
  import Label from "$lib/components/form/Label.svelte";

  export let id: string = crypto.randomUUID();
  export let label: string = '';
  export let name: string = '';
  export let options: [Entity];
  export let value: string | number | Array<string> | Array<number>;
  export let placeholder: string = '';
  export let disabled = false;
  export let required = false;
  export let tooltip = '';
  export let tooltipForLabel = '';
  export let cls = '';
  export let err = '';
  export let handleChange: (event: Event) => any = null;
  const dispatch = createEventDispatcher();

  $:isTip = !!err;
  $:isTipErr = !!err;
  $:tip = err;

</script>

<div
  class='form-control w-full mb-3 tooltip-top uppercase {cls}'
  class:tooltip={!!tooltip}
  data-tip={tooltip}
>
  <Label tooltip={tooltipForLabel} content={label} {id} uppercase/>

  <div class='tooltip-top' class:tooltip={isTip} class:tooltip-error={isTipErr} data-tip={tip}>
    <select
      class='select select-bordered w-full'
      class:select-error={isTipErr}
      id={id}
      {name}
      {disabled}
      {required}
      bind:value
      on:change={(e) =>{
        handleChange && handleChange(e);
        dispatch('change', e.target.value)
      }}
    >
      {#if placeholder}
        <option value='' selected>{placeholder}</option>
      {/if}
      {#each (options || []) as option}
        <option value={option.id}>
          {option.name}
          <svg width="400" height="110">
            <rect width="300" height="100" style="fill:rgb(0,0,255);stroke-width:3;stroke:rgb(0,0,0)"/>
          </svg>
        </option>
      {/each}
    </select>
  </div>
</div>

