<script lang='ts'>
  import {createEventDispatcher} from 'svelte';
  import IconSearch from '$lib/assets/svg/IconSearch.svelte';

  export let value = '';
  export let focus = true;
  export let cls = '';
  export let ref;

  const dispatch = createEventDispatcher();

  let id = crypto.randomUUID();
  let to;

  // debounce type input
  function watch(e) {
    if (to) clearTimeout(to);

    to = setTimeout(() => {
      value = e.target.value;
      dispatch('change', value);
    }, 600);
  }
</script>

<div class='relative {cls}'>
		<span class='absolute left mt-1.5 ml-2.5 '>
			<IconSearch/>
		</span>
  <input
    bind:this={ref}
    {id}
    type='text'
    placeholder='search...'
    class='input input-sm w-full pl-10 pr-10 relative bg-transparent focus:pl-8 focus:w-full transition-all ease-in-out duration-500'
    on:input={watch}
    value={value}
  />
</div>
