<script lang='ts'>
	import Input from '$lib/components/form/Input.svelte';

	export let label = '';
	export let name = '';
	export let placeholder = '';
	export let required = false;
	export let value = '';
	export let err = '';
	export let disabled = false;
  export let handleChange: (event: Event) => any = null;
</script>

<Input
	type='text'
	{label}
	{name}
	{placeholder}
	{required}
	{disabled}
  {err}
  {handleChange}
	bind:value={value}
	maskOptions={{ mask: '(000) 000-0000' }}
/>

