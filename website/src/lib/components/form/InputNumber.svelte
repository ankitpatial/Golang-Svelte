<script lang='ts'>
  import Input from '$lib/components/form/Input.svelte';

  export let label = '';
  export let name = '';
  export let placeholder = '';
  export let required = false;
  export let value;
  export let disabled = false;
  export let min = 0;
  export let max = 100;
  export let showRange = true;
  export let err ='';
  export let handleChange: (event: Event) => any = null;

</script>

<Input
  type='number'
  {label}
  {name}
  {placeholder}
  {required}
  {disabled}
  {err}
  {handleChange}
  bind:value={value}
  maskOptions={{
  mask: Number,
  min,
  max
}}
>
    <span slot='desc'>
      {#if showRange}
          between {min} - {max}
      {/if}
    </span>
</Input>

