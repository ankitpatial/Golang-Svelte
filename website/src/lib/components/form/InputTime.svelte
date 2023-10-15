<script lang='ts'>
  import Input from '$lib/components/form/Input.svelte';

  export let label = '';
  export let required = false;
  export let value = '00w 0d';
  export let disabled = false;
  export let w = 'w-full';
  export let handleChange: (event: Event) => any;

  // saveOnChange handler, will be used to invoke auto save
  export let saveOnChange: ((val: any) => Promise<void>) | undefined;
</script>

<Input
  type='text'
  {label}
  {required}
  {disabled}
  {w}
  {saveOnChange}
  {handleChange}
  bind:value={value}
  maskOptions={{
  	mask: '{W}w {D}d}',
		lazy: false,
		autofix: true,
		overwrite: true,
		placeholderChar: '0',
		blocks: {
			W: {
				mask: IMask.MaskedRange,
				from: 0,
				to: 52
			},
			D: {
				mask: IMask.MaskedRange,
				from: 0,
				to: 7
			}
  	}
	}}
>
  <span slot="desc">
    00-52 weeks 0-7 says
  </span>
</Input>

