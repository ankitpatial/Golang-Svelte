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
  export let w = 'w-full';
  export let m = 'mb-3';

  const dispatch = createEventDispatcher();

</script>

<div
  class='form-control {w} {m} tooltip-top uppercase'
  class:tooltip={!!tooltip}
  data-tip={tooltip}
>
  <Label content={label} {id} uppercase/>
  <select
    class='select select-bordered w-full '
    id={id}
    {name}
    {disabled}
    {required}
    bind:value
    multiple
    on:change={(e) =>{dispatch('change', e.target.value)}}
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

