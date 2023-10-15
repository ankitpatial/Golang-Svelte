<script lang='ts'>

  import Label from "$lib/components/form/Label.svelte";

  export let label = '';
  export let name = '';
  export let type: 'text' | 'email' | 'password' = 'text';
  export let value = '';
  export let cls = 'mb-0';

  let id: string = crypto.randomUUID();

</script>

<div class='form-control w-full'>
  <Label uppercase {id} content={label}/>

  {#if $$slots.group}
    <label class='input-group font-normal'>
      <input
        class='input input-ghost input-sm w-full  {cls}'
        id={id}
        {name}
        {type}
        {value}
        readonly
        disabled
      />
      <slot name='group'></slot>
    </label>
  {:else }
    <input
      class='input input-ghost input-sm w-full  {cls}'
      {id}
      {name}
      {type}
      {value}
      readonly
      disabled
    />
  {/if}
</div>

